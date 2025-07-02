## 一、请求处理流程（方法调用）

1. `TConndManager` 接收到客户端消息时，调用 `onMessage` 方法 将分发消息的任务提交$$到线程池中。（调用 `msgDispatcher.dispatch` 分发消息）
2. `msgDispatcher.dispatch` 调用 `MsgDispatcher` 类中的 `dispatch.handle.processClientMsg` 方法
3. `processClientMsg` 方法中通过 `baseMsgHandlerFactory.getMsgHandler(msg);` 调用当前 `BaseProtocolHandler` 实现类的`getMsgHandle` 方法，获取处理当前请求的 handle，最后调用当前handle中的 `handleClientRequest` 方法处理客户端请求
4. 其中 `GameMsgHandlerFactory implements BaseMsgHandlerFactory` 是处理游戏相关的工厂，它实现了 `getMsgHandler` 方法，在方法中通过**协议类型和反射**获取不同的处理器工厂，比如基于 SmartfoxServer 框架协议的 `HandlerRegisterCenter`，基于RPC 风格的 Protobuf 协议的`RpcPbMsgHandlerFactory` 以及基于普通 Protobuf 协议（CS通信）的 `PbMsgHandlerFactory`

### TconndMananger

`init` 方法中首先初始化的 `TconndApi` ，创建并启动了一个线程通过 `TcoondApi` 的本地方法专门向客户端发送消息，其中有三种将消息放入到 `sendMsgQueue` 队列的方法 1. 支持单播、无检查的 `sendMsgWithoutCheck` 2. 支持广播和单播、消息序列化后的 `sendMsgInternal`(广播session为null) 3. 支持组播的 `multCast` 基于上面上个方式扩展出兼容不同消息的方法 `sendMsg`。并且在 `init` 方法中创建了一个单线程的定时任务，专门用于监控 `TconndManager` 的运行状态。

| 方法             | 触发时机       | 核心功能      | 关键逻辑                                                        |
| -------------- | ---------- | --------- | ----------------------------------------------------------- |
| `onConnect`    | 客户端建立TCP连接 | 创建会话对象    | 1. 初始化Session对象 2. 存入`sessions`和`openIdSessions`            |
| `onMessage`    | 收到客户端数据包   | 协议解析+消息分发 | 1. Protobuf/SFS协议分流, `byte[]`封装为`BaseLogicMsg`对象 2. 提交线程池处理 |
| `onDisconnect` | 客户端断开连接    | 清理会话资源    | 1. 移除会话映射 2. 触发`onClientDisconnect`回调 • 处理日本服特殊逻辑           |

### SFS协议

登录请求是基于SFS协议，所以会通过 `MsgDispatcher` 经过一系列调用，最终调用 `GameMsgHandlerFactory` 中的 `getMsgHandle` 方法获取到SFS协议的请求处理工厂 `HandlerRegisterCenter`，`HandlerRegisterCenter` 中有 `getRequestHandler` 方法可以根据 **REQUEST_ID**到对应的 handle。继承 `AbstractClientRequestHandler` 实现 handleRequest 方法

### CS通信

在 `AbstractPbClientRequestHandler extends BaseClientRequestHandler` 中对实现了 `handleClientRequest` 方法，并在 `handleClientRequest` 中调用了`handle`，因此只需要继承 `AbstractPbClientRequestHandler` 并实现`handle` 即可

### RPC

继承 `AbstractRpcPbRequestHandler` 实现 handle 方法

## 一、登录

### 1. 登录逻辑

当客户端发送登录请求时，由 MsgDispatcher 在 `dispatch.handle.processClientMsg` 方法中获取 `LoginHandler` 处理器，并通过反射调用该处理器的 `handleClientRequest` 方法。

1. `handleClientRequest` 首先会执行检查游戏是否正在关闭，用户是否在黑名单，登录监控等操作，然后调用 `handleRequest`方法

2. `handleRequest` 中会检查服务状态，将登录信息封装成LoginInfo，然后根据openid对同一账户加锁，防止并发登录，然后调用 `handleInternalRequest` 方法

3. `handleInternalRequest` 中获取服务器id，分区id等，然后通过 `userSession.sendMsg` 把封装好的基础信息发送给客户端，最后调用 `handleZoneJoinEvent`

4. `handleZoneJoinEvent` 中通过 `UserService` 加载用户档案，通过 `userProfile.toLoginInfo()`获取封装后的用户信息，并通过 `GameEngine.getInstance().pushMsg` 将封装信息发送给客户端，将用户加入到游戏区服

5. `userProfile.afterLogin` 登录后通过这个方法加载玩家的联盟信息、VIP状态、技能、装备等游戏数据，发送心悦、回流等活动。其中也调用了 `loginHelper` 中的`afterLoginNtf`方法

6. `afterLoginNtf` 方法将一些其他消息通知通过异步的方式发送给客户端

### 2. 过程中用到的Java锁：

1. synchronized  (userMsgSync) ，TconndManager中在客户端发来消息时，通过synchronized确保每个用户的多条消息安全放入队列中。加锁方式：修改加锁对象的头信息，确保当前线程才可以获取
2. AtomicInteger 原子锁，实现原子的增减  。加锁方式：cpu级别的优化
3. 自定义 `ReentrantMTLock`，并通过 `SafeLockMgr` 管理，实现锁可以通过定时任务线程来释放，而非必须是持有锁的线程   

### 3. 涉及的mysql

1. ` userProfile = UserProfile.getWithUidFromDb(uid);` 当缓存中没有 `UserProfile` 会从db中获取 `UserProfile` 

```sql
<select id="selectByUID" resultMap="UserProfile" parameterType="java.lang.String">
    SELECT * 
    FROM userprofile
    WHERE uid = #{uid,jdbcType=VARCHAR} order by regTime desc limit 1
</select>
```

2. `userProfileAttr = mapper.getWithUid(uid);` 获取 `UserProfileAttr` 

```sql
<select id="getWithUid" resultMap="BaseResultMap">
    select
    <include refid="Base_Column_List" />
    from userprofile_attr
    where uid = #{uid}
</select>
```

3. `userProfile.updateDeviceToken(loginInfo.getDeviceToken());`   更新设备令牌

```sql
UPDATE userprofile
SET
    deviceToken = #{deviceToken,jdbcType=VARCHAR},
    deviceTokenUpdateTime = #{updateTime,jdbcType=BIGINT}
WHERE uid=#{uid,jdbcType=VARCHAR}
```

4. `shieldArr = getShieldDetailList(uid);` 查询用户聊天屏蔽列表

```sql
SELECT 
    cs.uuid, cs.other, u.serverId, u.name, u.headImg, u.locate, p.power, a.abbr 
FROM 
    chat_shield cs
LEFT JOIN 
    userprofile u ON cs.other = u.uid
LEFT JOIN 
    playerinfo p ON cs.other = p.uid
LEFT JOIN 
    alliance a ON u.allianceId = a.uid 
WHERE 
    cs.owner = ?
```

5. `PrayService.getLoginInfo(this, outData);` 中查询祈祷的相关信息

```sql
SELECT * FROM user_pray WHERE uid = ?
```

6. `List<AllianceScience> allianceSciences = mapper.selectAllianceAllSciences(allianceId);` 获取联盟科技相关信息

```sql
SELECT * FROM alliance_science
    WHERE allianceId = #{allianceid,jdbcType=VARCHAR}
```

7. `GameService.fillServerTime(retObj);` 获取数据库服务器的当前 UTC 时间

```sql
ISFSArray utcArray = SFSMysql.getInstance().query(
    "SELECT unix_timestamp() AS db_utc_timestamp, " +
    "EXTRACT(HOUR FROM TIMEDIFF(NOW(), UTC_TIMESTAMP())) AS offset"
);
```

8. `resetAllianceTaskInfo();//联盟宝藏每日次数刷新` 查询联盟任务完成记录

```sql
<select id="selectByPrimaryKey" resultMap="BaseResultMap" parameterType="java.lang.String" >
        select
        <include refid="Base_Column_List" />
        from alliance_task_record
        where uid = #{uid,jdbcType=VARCHAR}
    </select>
```

9. `handleZoneJoinEvent` 中通过jdbc更新最后在线时间

```sql
UserService.updateLastOnlineTimeWithCache(userProfile);
String executeSql = "update account_new set  lastOnlineTime = ? where gameUid = ?";
```

### 4. Redis操作

`RedisSession` 通过连接池管理工具类 `RedisSessionUtil.getInstance()` 来管理Jedis不同连接模式，通过 RedisSession 的静态方法 `getGlobal`、`getShare`来创建RedisSession，

1. `rs.set(key, String.valueOf(1));` 登录时标记为关闭界面
   
   ```sql
   String key = "HONGMENG_SCIENCE_PAGE_OPEN" + uid;
   ```

2. 在 `afterLogin` 中对判断用户是否是回流玩家，如果是回流玩家，则将用户openid放入到全局和共享缓存中

```java
public void addLogin() {
    if (isHuiliuUserByCache(userProfile)) {
        long activityBeginTime = getActivityTime().getFirst();
        if (activityBeginTime == 0) {
            return;
        }
        String key = REDIS_KEY_IS_REBACK_LOGiN + activityBeginTime;
        RedisSession globalRedis = RedisSession.getGlobal();
        globalRedis.sAdd(key, userProfile.getOpenid());
        globalRedis.setPExpireTime(key, DateUtils.getNewDay());

        RedisSession shareRedisSession = RedisSession.getShareRedisSession();
        shareRedisSession.sAdd(key, userProfile.getOpenid());
        shareRedisSession.setPExpireTime(key, DateUtils.getNewDay());

        WecLoggerUtil
                .logRebackManagerFlow(userProfile, userProfile.getUid(), "addLogin", userProfile.getOpenid(),
                        1);
    }
}
```

3. 在 `afterLogin` 中对距离上次登录间隔24小时的用户，删除全局缓存中的openid

```java
if(System.currentTimeMillis()-userProfile.getOffLineTime() > 24 * 3600 * 1000) {
        RedisSession globalRedis = RedisSession.getGlobal(true);
        globalRedis.hDel("million_R_Lose", userProfile.getOpenid());
        logger.warn("uid:{} BigR remove from million_R_Lose redis cache,lastOnlineTime:{},lastOfflineTime:{},cur:{}",
                userProfile.getUid(),userProfile.getLastOnlineTime(),userProfile.getOffLineTime(),System.currentTimeMillis());
        AttrApi.add(35478573    ,1);
    }
```

### 5. Tlog流水表

日志通过对应java类的 **`logToTlogd()`** 方法通过udp协议发送到日志服务器

1. 在 `handleLogin` 中对于新用户的分支中记录了邀请注册流程的日志，如邀请人uid、被邀请人uid等信息

```java
WecLoggerUtil.logInviteRegisterFlow(userProfile,
                                InviteRegisterManager.InviteRegisterType.JuyiInviteRegister.ordinal(),
                                1, userProfile.getOpenid(), inviteData.getExtra(), inviteData.getUid(),
                                userProfile.getUid());
```

2. 在 `handleLogin` 中记录了军队相关操作流水

```java
WecLoggerUtil.logArmyFlow(userProfile, TlogMacrodef.ARMYOPTYPE.LOGIN, "0", 0, 0, 0, 0, 0, 0, 0, "",
                userProfile.getArmyManager().getAllArmyStr());
```

3. 在调用完 `UserProfile.afterLogin` 后会记录玩家或系统触发的奖励返还操作流水

```java
WecLoggerUtil.logRebackManagerFlow(userProfile, userProfile.getUid(), "addLogin", userProfile.getOpenid(), 1);
```

4. 在调用完 `UserProfile.afterLogin` 后会记录邀请好友发送奖励的流程日志

```java
WecLoggerUtil.logRebackSameSvrSNSFlow(userProfile, inviterUid, isRegister);
```

5. 在调用完 `UserProfile.afterLogin` 后记录 superLottery 日志

```java
WecLoggerUtil.logSuperLotteryThemeFlow(userProfile, themeIdOld, superLottery.getThemePoolId(), 0);
```

6. 记录游戏内战友绑定/解绑关系的流水日志

```
 // Tlog
            WecLoggerUtil.logTlogComradeBindFlow(userProfile, comradeUid, uid, TlogMacrodef.ComradeBindOpType.Comrade_NaturalUnbind, 0);
            WecLoggerUtil.logTlogComradeBindFlow(null, uid, comradeUid, TlogMacrodef.ComradeBindOpType.Comrade_NaturalUnbind, 0);
```

7. 在调用完 `UserProfile.afterLogin` 后记录战阵属性变化日志

```
WecLoggerUtil3.logSwordWarHistoryAttrFlow(userProfile, 1, preFormationInfo.getArmyTypeSign(), attrType,
                                    (float) maxAttrNow, (float) newAttr);
```

## 二、活动

### 活动配置的加载

1. 在 `ActivityConfig` 的 `initalize` 方法中调用 `refreshAllXmlFile` 方法获取所有满足正则表达式 `"^activity_center.*\\.xml$"` 文件名，后续调用`execRefresh.updateConfigMap.loadActivityConfigloadActivityConfigFromXml` 方法解析加载对应 xml 配置文件，并且在 `updateConfigMap` 中会查询 db，合并xml活动配置和db活动配置，并把所有活动配置保存在 `configMap` 中，ActivityConfig中的 `configMap` 字段存储的是xml和db中所有的活动，包含过期的活动。

2. `ActivityHandle` 是活动配置的handle基类，包含活动xml配置的信息。

3. `ActivityProxy` 管理并初始化`ActivityConfig` 实例，并且在ActivityConfig的`onConfigRefresh()` 中对 `ActivityProxy` 的 `handlerMap` 字段进行初始化，`handlerMap` 中存储时所有没有过期的活动配置。(通过`ActivityHandlerFactory.newActivityHandler()` 方法创建活动配置handle的实例，存储在 `handleMap` 中)

4. `ActivityHandlerFactory` 提供静态方法来创建 `ActivityHandle` 实例

5. `ActivityCenter` 用来管理活动，包含全量通知和增量通知客户端，会在登录时通过`UserProfile`初始化时(toLoginInfo)，通过ActivityCenter的 `initialize` 中初始化所有符合条件的活动。

在 `GameScheduler` 中会启动一个固定的线程来调用 `ActivityProxy` 的 `execRefresh` 来根据配置文件的修改修改内存中的活动配置

### 活动配置的函数调用

GameSvrApp会在main方法中调用 init() 方法，init方法中有 `GameEngine.getInstance().init();` ，gameEngine的init() 方法会调用 `initGameData()` ，其中包含 `ActivityProxy.getInstance().init()` ，activityProxy 的init() 方法包含 `activityConfig.initialize();` ，最终完成对 `ActivityConfig` 和 `ActivityProxy` 的初始化，加载xml和db活动配置加载到内存中。并且会有单独的定时线程监控xml配置文件，当配置文件新增活动时会加载到内存。（GameScheduler的定时线程是在gameEngine.init()中启动的）

`ActivityHandler` 中有 `getXmlConfig()` 方法，获取当前活动的xml文件中的配置信息。`ActivityConfig` 和 `ActivityProxy` 在项目启动时完成初始化，每个运行时活动的xml配置在项目启动时初始化完成，并存放在ActivityProxy的`handlerNameMap`字段中，key为活动配置的java类名，`ActivityCenter` 在玩家登陆时给对应的玩家进行初始化

活动handle开发

活动客户端请求处理需继承`AbstractPbClientRequestHandler`方法，实现其中handle，可以通过PbResLoader、PbResLoader2、PbResLoader3获取活动的具体配置信息。（主要是活动excel配置信息转成pbin后，加载到PbResLoader中的消息结构中）

## 三、国宝

``NationalTreasureMgr`` ：管理国宝的获取、升级、过期、奖励发放等流程。

- 处理国宝的获取、合成、过期处理
- 管理国宝抽奖功能，如获取、发送、重置玩家抽奖信息，将国宝添加到属性系统等
- 加载和管理国宝相关配置

`NationalTreasureEffectUtils` ：计算国宝的具体加成数值，不管理流程。

- 计算国宝的基础属性加成​​、国宝套装的效果、国宝之间的相互加成

- 特殊效果加成（神兽，联盟成员共享）

`NationalTreasureLottery` 是一个 ​用户扩展属性attr，包含用户抽奖相关的信息。在UserAttr中被引用，UserAttr会序列化成二进制存储在 UserProfileAttr 的 `byteFull` 字节数组中

### 1. 游戏国宝静态配置加载

主要通过 `PbResLoader3` 工具类的`setAllServerRes3`方法设置 ResMapAllServer3字段的值。ResMapAllServer3 是游戏的配置类（protobuf生成的）。包含了国宝的配置

在游戏后台启动时，GameSvrApp的init()方法中，通过GameEngine的init()方法，调用**GameConfigManager的`initGameSvrConfig` 方法**。在这个方法中会通过 GameConfigLoader的loadpb方法`PbResLoader.load()` 其中会通过 `PbResLoader3.setAllServerRes3(tmpRes3.build());` 完成从 Protobuf 二进制文件中加载国宝的配置内容到protobuf结构中

（`GameConfigManager.initGameSvrConfig(getServerOpenTime());` 主要用来加载游戏的 xml 和 pb配置数据）











### 2. 抽国宝

客户端的抽奖请求由 `NationalTreasureLotteryMsgHandler` 处理。首先通过 uid获取当前的UserProfile和客户端请求的消息，接着判断是否可以抽奖，如果可以则抽奖，并把抽奖结果通知给客户端，抽奖得到碎片可能可以合成，所以通过 `NationalThreasureMgr` 中提供的国宝合成方法，如果可以合成则合成国宝。

### 3. 国宝合成

国宝合成由 `NationalTreasureMgr` 提供对应的成员方法 `doTreasureCompose` 进行国宝的合成。首先确保国宝系统开放和城堡等级达标，接着从ResMapAllServer3中获取国宝的配置，构建不同品质所需碎片数量的映射。然后获取用户解锁的国宝，然后遍历所有解锁后的国宝分别对红色品质的国宝和非红的进行合成。并通过`addTreasureToUserAttr` 将国宝信息更新到属性系统中，最后将合成的国宝信息发送给客户端

### 4. 国宝升级，升星

升级：国宝升级由 `NationalTreasureLevelUpMsgHandler` 处理，获取当前用户和客户端发送消息中要升级的国宝id，并校验是否符合要求（如是否解锁该国宝，是否满级等），对于满足要求的国宝升级请求，通过 ItemManager 减去耗材，并设置国宝等级，最后通过`NationalTreasureEffectUtils.onTreasureEffectUpdate(userProfile);` 刷新国宝效果（addEffect可以使属性生效），并通知客户端。

升星：国宝升星由 `NationalTreasureStarUpMsgHandler` 处理，获取当前用户和客户端发送消息中要升级的国宝id，并校验是否符合要求（如是否解锁该国宝，是否满星等），对于满足要求的国宝升星请求，通过 ItemManager 减去耗材，并设置国宝星级，最后通过`NationalTreasureEffectUtils.onTreasureEffectUpdate(userProfile);` 刷新国宝效果（addEffect可以使属性生效），并通知客户端

### 5. 国宝的效果生效

国宝生效主要通过 GameEffectMgr 中的 `onNationalTreasureEffectChange` 方法，其中会先使所有国宝的效果都失效，然后通过 addEffectImp 使国宝效果生效。（国宝加成，国宝套装加成，国宝收集加成等）

- 1. 在登录加载时会通过 `onNationalTreasureEffectChange`使国宝生效

- 2. 在登录成功后(stdMgrAfterLoginSucc)，客户端会发送一个login.success协议的请求，并由 `LoginSuccHandler` 处理，调用GameMgr 的 `stdMgrAfterLoginSucc`方法，该方法中会调用 BaseGameMgr 实现类的 `stdMgrAfterLoginSucc` 方法。其中 NationalTreasureMgr 实现了 `BaseGameMgr` 所以国宝的 `stdMgrAfterLoginSucc` 方法会被调用，该方法中会  刷新过期国宝，生效国宝效果等。(通过 UserProfile 获取 GameEffectMgr，然后调用`onNationalTreasureEffectChange` 方法)

- 3. 在客户端发送刷新国宝请求时会移除超时的国宝，`NationalTreasureRefreshMsgHandler` 会调用 `NationalTreasureMgr` 的`refreshTimeOutNationalTreasure()` 方法，重新使国宝生效

`GameEffectMgr` 用来管理游戏效果，其中的addEffectImp用来将游戏效果加入到属性系统中， `GameEffects` 中存储着多个不同种类的游戏效果，可以通过 效果ID 查询到对应的 `GameEffect` ，`GameEffect` 中存储着当前effectId的值 `GameEffectValue`

6. 收集进度

收集进度主要在NationalTreasureMgr中的`calculateTotalScore`处理，该方法会根据玩家收集国宝的品质，星级，等级情况计算一个收集进度数值，并通过 sendTreasureInfoNtf 返回给客户端

7. 关键tlog

`NationalTreasureLevelUpMsgHandler` 中`WecLoggerUtil3.logNationalTreasureLevelUp(userProfile, treasure.getId(), level, itemIdBuffer.toString(), itemCountBuffer.toString());` 记录国宝升级 tlog

`NationalTreasureStarUpMsgHandler` 中

`WecLoggerUtil3.logNationalTreasureStarUp(userProfile, treasure.getId(), star, pieceId, pieceNum, starUpConf.getCostItemId(), starUpConf.getCostItemNum());` 记录国宝升星 tlog

`NationalTreasureMgr` 中

`WecLoggerUtil3.logNationalTreasurePieceExchange(userProfile, pieceId, pieceNum, nationalTreasurePiece.getCount());`记录国宝碎片的变更情况

`WecLoggerUtil3.logNationalTreasureComposeFlow(userProfile, treasureConf.getId(), treasureConf.getPieceId(), needNum, flushTime);` 记录玩家合成非红色品质国宝的操作

`WecLoggerUtil3.logNationalTreasureComposeFlow(userProfile, treasureConf.getId(), treasureConf.getPieceId(), needNum, flushTime);`记录玩家合成红色品质国宝的操作

`WecLoggerUtil3.logNationalTreasureRewardFlow(userProfile, treasure.getId(), treasure.getStar(), treasure.getLevel(),  
        TlogMacrodef.NationalTreasureRewardType.REWARD, treasureConf.getRewardId(), rewardVal, 0, 0); `记录奖励发放的具体内容



### 6. 根据excel生成所需配置

1. 在 `trunk\common\excel\proto` 中创建对应的 .proto 文件，文件名为对应 excel 文件中的，如国宝excel配置中 `convert(ResNationalTreasure.proto, table_NationalTreasureStarUp, NationalTreasureStarUpConf.pbin)` ，则对应的文件名是 `ResNationalTreasure.proto`
2. 根据 excel 中的字段，在 .proto 文件中配置对应的 枚举 或 message，如国宝excel配置中国宝品质有5种，在 `ResnationalTreasure.proto` 中定义了枚举 ENationalTreasureQualityType ，并设置了5中枚举值
3. 在 `trunk\common\excel\proto\ResMapAllServer3.proto` 中添加该 .proto 文件的map引用字段，如 `map<int32, NationalTreasure> NationalTreasureConf = 3284;`

4. 执行`trunk\common\excel\ClientExcelConverter.exe`，将excel表数据转成对应 **.pbin** 和 **.txt** 文件
5. 执行`trunk\common\genProtoCode.bat`，生成 .proto 文件对应的 java 文件，如国宝的 `ResNationalTreasure.proto` 生成对应的 `ResNationalTreasure.java` ，该类中包含 get、set、newBuilder()、序列化，反序列化等方法

注意事项：

创建 .proto 和 excel 文件需要 svn add

table_XXX中的字段格式务必为 repeated Xxx rows



### 7. 配置加载

excel表，游戏的静态配置或数据 `\common\excel\xls` 

proto 文件，定义pb结构，用于存储 **.pbin** 中的数据 `\common\excel\proto`



`ResNationalTreasure.proto` 根据excel中的字段定义了所有国宝相关的pb表结构，如国宝的配置：国宝品质 ENationalTreasureQualityType、国宝分类 ENationalTreasureType、国宝效果 NationalTreasureEffect、国宝 NationalTreasure

然后通过定义`message table_NationalTreasure + repeated`定义国宝的列表，之后将 `NationalTreasure` 的所有 **NationalTreasureConf.pbin** 内容加载到 `table_NationalTreasure` 中。（在loadOneField方法中实现时，也有可能加载.txt文件，当.txt最新修改时间大于60s时）

![image-20250630150727542](D:\notes\笔记图片\image-20250630150727542.png)



在 `ResMapAllServer3` 中定义了 message ResMapAllServer3，并用map结构引用了 **NationalTreasure**、**NationalTreasureSet**等国宝表结构，对应的字段名分别是 **NationalTreasureConf**、**NationalTreasureSetConf** 等。

在项目启动时会把 **table_NationalTreasure**、**table_NationalTreasureSet** 等列表中的值插入到对应的 map 中，k为id，val为对应的行值。（这个操作是在 `PbResLoader` 的 `loadRowValues()` 方法中通过反射实现的）

![image-20250630152115137](D:\notes\笔记图片\image-20250630152115137.png)















## 四、属性系统

`UserProfileAttr` 将存储了 `UserAttr` 的变量，序列化为二进制内容，存储在对应表中的 ` byteFull` 字段。` UserAttr` 中的对象并没有对应的表，数据是以二进制的形式存储在了UserProfileAttr表中的 bytefull 字段。

其中UserAttr中除了包含uesr相关的基本属性，还包含有扩展属性 `ExtendAttr`。（`ExtendAttr` 中包含有用户国宝相关的信息 `NationalTreasure`）

用户effect相关的属性，都由GameEffects基于属性系统 AttrObj 管理。对于某个用户属性的根节点类，并且会为该类创建初始化脏数据标记数组 byte[] dirtyBytes（为1表示该对象这个位上对应的属性是脏的）。对于GameEffects中的AtrrObj变量，他的owner会指向父类，root会指向根对象，当出现脏数据时会递归向上标记脏数据。

属性系统的回写：

通过 `UserProfileAttr` 的 `saveAttr2DbLogic` 方法，其中会通过 `UserAttr` 的 `userAttrFull.copyToDto()` 将全量数据复制到Protobuf结构，并通过 `AttrDbUserAttr.proto_db_UserAttr.newBuilder().build().toByteArray()`  序列化字节数组，最后通过 `updateDbDirectly(true);` (true表示全量写入db，false增量)方法中调用mybatis将bytefull或bytechange写入数据库



### 1. 向属性系统中添加属性

1. `common\attr\vc.attr` 中的根节点 **@UserAttr**，对应代码中的 UserProfileAttr.java中的 userAttrFull成员变量

+onlydb： 字段只在gamesvr内存和db中有，不下发客户端

+onlycs: 字段只在gamesvr内存和客户端有，不存储db，客户端每次重登时会重新初始化为默认值



2. 通过 `\trunk\common\attr\genattr.bat` 生成对应的 **.proto** 和 **java** 文件

`\trunk\common\cs_proto\proto\attr*.proto` 中存储了属性系统的**cs**客户端和服务器交互协议，

`\trunk\common\db_proto\attr*.proto` 中存储了属性系统的序列化到**db存储**的协议文件

例如：`vc.attr` 中的

```
#国宝碎片
@NationalTreasurePiece
id         int32                    1 #碎片ID
count      int32                    2 #碎片数量
addTime    int32                    3 #碎片增加的时间,单位天
```

生成的对应 .proto 文件是

```protobuf
syntax = "proto2";
option cc_generic_services = false;
package com.tencent.wc.protocol;

message proto_db_NationalTreasurePiece {
	optional sint32 count = 2;
	optional sint32 addTime = 3;
	optional sint32 id = 1;
}
```

生成的对应attr的 .java 文件是

```java
public class NationalTreasurePiece extends AttrObj {
    private int count;
    private int addTime;
    private int id;

    private final static int ATTR_INDEX_COUNT = 0;
    private final static int ATTR_INDEX_ADDTIME = 1;
    private final static int ATTR_INDEX_ID = 2;
    private final static int _ALL_ATTR_COUNT = 3;

    public NationalTreasurePiece() {
        super(_ALL_ATTR_COUNT);
		count = 0;
		addTime = 0;
		id = 0;

        setOwner(null, -1);
        flushDirty(false);
        init = true;
    }
...
```



3. 将proto转成 java文件 `common\genProtoCode.bat` 根据 .proto文件 生成 protobuf java文件，比如：

db protobu 数据表

```java
package com.tencent.wc.protocol;

public final class AttrDbNationalTreasureLottery {
  private AttrDbNationalTreasureLottery() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }
...
```

cs protobuf 数据表

```java
package com.tencent.wc.protocol;

public final class AttrNationalTreasureLottery {
  private AttrNationalTreasureLottery() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }
...
```



## 五、武将委任

### 1. 配置

1. Res 配置

建筑委任：GeneralAppointData

一键委任：GeneralOneButtonAppointConf

进阶委任：GeneralUpLevelAppointConf

武将分类：GeneralFenLeiConf

委任建筑：GeneralAppointBuildListConf

技能对照：NewGeneralSkillData

2. cs 协议 配置

```prot
//新委任协议
message GeneralAppointSingle {
    optional int32               index			          = 1; //
	optional int32               genId                    = 2;
	optional int32               page                     = 3;
}

message GeneralAppointReq {
    optional int32               buildingId			        = 1; //
	repeated GeneralAppointSingle            appoint        = 2;
}

message GeneralAppoint_C2S_Msg {
    optional int32               oneKey			        = 1; //
	repeated GeneralAppointReq   appointReqs            = 2; //
}
```



3. handle

武将的委任，取消委任，一键委任的 `GeneralAppoint_C2S_Msg` 消息请求，都会由 `GeneralAppointMsgHandler` 处理

其中将一个武将委派到某个建筑的主要逻辑是在 `grantManyAppoint(ISFSObject appointData, int oneKey)` 中的，一键委任的逻辑是在`grantAppointOneButton()` 中的

