学习 Vue.js 应该像学习一门编程语言一样，首先要熟练掌握常用的知识，而对于不常用的内容可以简单了解一下。先对整个框架和语言有一个大致的轮廓，然后再逐步补充细节。千万不要像学习算法那样，一开始就钻牛角尖。

## 前序:

vueAPI的风格分为: 选项式和组合式，vue2中一般用选项式, 所以文章中讲到的api一般都是组合式

vue3官方推荐用基于 Vite 的构建项目，首先安装 [node.js](https://nodejs.org/zh-cn), 然后在终端输入
`npm create vue@latest`，根据提示构建项目。`npm install` 和 `npm run dev` 分别是下载依赖和运行项目

当然你也可以通过`vue ui`来构建项目

## 一、初识vue：

### 1. 项目结构:

当我们新建一个vue项目时, 整个项目刚开始会有很多`.vue, .js, .html .css`文件，但真正有用的只有`index.html, App.vue和一些配置文件`

vue组件包含三个区域, 分别用来写`html,js(ts),css`  其中`style`的`scoped`可以让css样式的作用域限定到当前组件
```vue
<template>
</template>
<script>
</script>
<style scoped>
</style>
```

### 2. vue常用函数

`createApp():`是Vue 3中用来创建一个新的Vue应用实例的方法。它接受一个根组件作为参数，并返回一个应用实例。根组件就是export出去的name字符串

`mount():` 是把createApp创建的应用实例挂载到一个容器上。 接受的参数容器可以是一个css的选择器

```vue
createApp(App).mount('#app') // App.vue文件中
<div id="app"></div>	//index.html中
我们刚创建的项目中会有上面的两个文件, 实际上就是创建了一个App实列, 通过id选择器把它挂载到一个div上
```


`setup():` 是一个函数, 函数里面可以定义响应式变量、计算属性、方法, 然后把变量, 方法 `return` 出去, 就可以在该组件的`template`中使用。 （我们一个组件所用到的函数、变量都写在这里面）

我们每个组件的`name、setup、components...`都需要export出去。 其中`name`是这个组件的名字, 当其他组件要用到当前组件的时候就是要`<name/>`, 

如下: 在`setup`函数中定义了**name**和**click**方法给当前组件的html使用。 代码中的(`@用来绑定事件`, 包括html所有事件和自定义事件)


```vue
<template>
  <h1>{{ age }}</h1> 		<!-- {{ }} 双大括号将 Vue 实例中的数据绑定到 HTML 元素中 -->
  <button @click="handleClick">click</button>
</template>

<script lang="ts">
export default {
  name : "App",
  setup : () => {
	    let age= 10;
	    const handleClick = () => {
	    age = 18;
    }
    return {name, handleClick}
  }
}
</script>
```
前面讲到setup里面定义的方法，变量等都需要`return`出去之后 `html` 中才能使用，下面这个语法糖可以让我们不用一个个写`return`

`setup语法糖`: 在`<script setup>  </script>`中写方法、状态的话，不需要把方法或状态return回去, vue会自动帮我们`return`, 而且也不需要用`components`注册用到的组件。

`ref():` 是一个函数，使用 ref 函数可以创建一个包装过的响应式对象，使其能够在 Vue 组件中进行响应式数据绑定。ref接受的参数是**基本类型、数组、对象类型**的。(当参数是**对象**时, ref内部其实是调用的reactive)

`reactive()`: 使用 reactive函数可以创建一个包装过的响应式对象，reactive接受的参数是**对象类型**的

**响应式数据**指的是当我们的变量发生变化的时候, 跟这个变量相关的元素就会**自动刷新**。
上面的代码中我们点击button的时候会把age = 18, 但实际上我们页面显示的时候age还是等于10, 因为age不是响应式的变量
所以我们需要用ref创建age变量, `let age = ref(10)`, 这样age发生变化时, 跟age相关的元素也会自动刷新。(需要**注意:** 在js中使用由ref创建的响应式对象是要加上".value", 比如 age.value；在html中则不需要, 可直接写成age)

`toRefs`和`toRef`:当我们把一个对象或者变量解构的时候。假设`person`对象中有`name, age` `const {name, age} = person`，那么此时解构出来的的`name`和`age`不再具有响应式。为了使解构出来的仍然具有响应式,那么应该用`toRefs`或`toRef`包括起来。 `const {name, age} = toRefs(person)`.

`computed()`:是基于其他数据的值计算出来属性，并且当这些数据变化时，计算属性的值会自动更新

计算属性是基于它们的依赖进行缓存的，只在其依赖发生变化时才会重新计算，而方法每次调用时都会重新执行。

如下, 如果是基于函数`getFullName()`计算全名的话当name发生变化是控制台会输出3次1, 如果是用`FullNameByComputed`的话, 控制台只会输出一个1
```vue
<span>{{ getFullName() }}</span><br>
<span>{{ getFullName() }}</span><br>
<span>{{ getFullName() }}</span><br>

setup : () => {
    let firstname = ref('');
    let lastname = ref('');

    const getFullName = () => {
      console.log(1); 
      return firstname.value + lastname.value;
    }
	const FullNameByComputed= computed(() => {
      console.log(1); 
      return firstname.value + lastname.value;
    });
    
	return {firstname, lastname, getFullName, FullNameByComputed}
}
```

通过上面的computed创建的`FullNameByComputed`是一个类似ref创建的对象, 在`js(script标签)`中要访问`FullNameByComputed`的话需要`FullNameByComputed.value`。
如果需要直接修改value的话需要在创建对象的时候传递get和set方法，如下

```vue
const getFullNameComputed = computed({
      get: ()=> {
        return firstname.value + "-" + lastname.value;
      },
      set: (newName)=> {
        let name = newName.split('-');
        firstname.value = name[0];
        lastname.value = name[1];
      }
    });
```
上面代码中，当我们`FullNameByComputed.value`的话, vue会自动调用get方法, 当我们FullNameByComputed.value=newName的时候，会自动调用set方法, 并把等号右边的值传递到set参数中

### 3. 常用语法糖

`v-model`: 是 Vue.js 中用于在表单控件（如输入框、复选框、单选按钮等）和组件之间创建双向数据绑定
如下: 我们让`username`和输入框双向绑定, 当用户在输入框中输入字符时, username 就会变成用户输入的值。其中`username`需要是响应式的变量

```vue
<input v-model="username" type="text">
<p>{{ username }}</p>
let username = ref('');
```
`v-for`: 通过 `... in ...` 用来遍历数组
如下: 假设我们有一个persons数组, 数组长度是3，那么我们可以通过`v-for`来生成3个`li`标签。其中对于遍历生成的标签我们需要通过`:key="“`给他绑定一个唯一的key, 一般用id。(注意key前面的`:`, 不写`:`的话，vue会把person.id当成字符串解析)
```vue
<li v-for="person in persons" :key="person.id">{{ person.name }}</li>
```
`v-if`和`v-else`就如他们的字面意思, 判断用的

注意: 上面文章和下面文章中用到的函数需要自己引入，如`computed`、`ref`等等，文章中就不再过多赘述了。

## 二、路由

### 1. 基础概念

**路由（Routing）**：在单页应用（SPA）中，路由用于管理不同视图或页面之间的导航。用户点击链接时，路由会根据 URL 的变化加载相应的组件，而不需要重新加载整个页面。

**Vue Router**：Vue Router 是 Vue.js 的官方路由管理器，提供了灵活的路由功能，通过定义路由规则来管理组件的展示。

### 2. 常用组件

Vue Router 提供了一些核心组件，以下是常用的组件及其用途：

- **RouterView**：用于渲染匹配的路由组件。它通常在根组件中使用。

  ```html
  <template>
    <RouterView />
  </template>
  ```

- **RouterLink**：用于在应用中创建导航链接的组件，类似于 HTML 的 `<a>` 标签。

  ```vue
  <template>
    <RouterLink to="/about">About</RouterLink>
  </template>
  ```

**to的三种写法**
`<router-link to="">点击跳转</router-link>`中的`to`有三种写法.(`to`前面加上`:`, 才能让vue不把后面的内容解析成字符串)
1. `to="/home"`
2. `:to="{path : "/home"}"`
3. `:to="{name: "HomeView"}"`

### 3. 路由配置

配置路由涉及创建一个路由表，定义各个路由及其对应的组件。以下是一个简单的示例，展示如何在 Vue 3 应用中配置路由。

```javascript
// router/index.js
import { createRouter, createWebHistory } from 'vue-router';
import Home from '../components/Home.vue';
import About from '../components/About.vue';
import NotFound from '../components/NotFound.vue';

// 路由模式
// 1. Hash 模式路由使用 URL 的 hash（#）部分来模拟一个完整的 URL，实际 URL 的路径部分始终是 # 符号之前的内容
// 2. History 模式利用 HTML5 History API 来管理路由，这使得 URL 看起来像普通的 URL，没有 # 符号。

const router = createRouter({
  history: createWebHistory(),
  routes: [
  {
    path: '/',
    component: Home
  },
  {
    path: '/about',
    component: About
  },
  {
      path: '/:pathMatch(.*)',	// 匹配所有路径
      redirect: '/404',	//重定向到/404
   },
  {
    path: '/:catchAll(.*)', // 捕获所有未匹配的路径
    component: NotFound
  }，
];

});

export default router;
```

### 4. 嵌套路由

嵌套路由允许你在某个父路由的组件内部定义子路由。以下是一个示例：

```js
// router/index.js
const routes = [
  {
    path: '/users',
    component: UserProfile, // 父组件
    children: [
       {
         path: "",
         redirect: "/users/home/all", // 默认的 /users 会重定向到 /users/home/all
       },
      {
        path: 'settings',
        component: UserSettings // 子路由
      },
      {
        path: 'activity',
        component: UserActivity // 子路由
      }
    ]
  }
];
```

在 Vue 中，可以使用 `RouterView` 作为子组件的容器，允许在父组件中渲染子路由：

```html
// App.vue
<template>
  <RouterView />
</template>

<script setup>
// 这里可以添加相关逻辑
</script>
```

在父组件中定义嵌套路由：

```html
// UserProfile.vue
<template>
  <div>
    <h1>User Profile</h1>
    <nav>
      <RouterLink to="settings">Settings</RouterLink>
      <RouterLink to="activity">Activity</RouterLink>
    </nav>
    <RouterView /> <!-- 渲染子路由 -->
  </div>
</template>

<script setup>
// 这里可以添加 UserProfile 组件的逻辑
</script>
```

### 5. RouterView 和 RouterLink 的使用

在父组件中使用 `RouterLink` 导航到子路由，并通过 `RouterView` 渲染匹配的子组件。

```html
<!-- UserProfile.vue -->
<template>
  <div>
    <h1>User Profile</h1>
    <nav>
      <RouterLink to="settings">Settings</RouterLink>
      <RouterLink to="activity">Activity</RouterLink>
    </nav>
    <RouterView /> <!-- 渲染匹配的子路由 -->
  </div>
</template>

<script setup>
// 这里可以添加 UserProfile 组件的逻辑
</script>
```

### 6. 动态路由

动态路由允许在路径中使用参数，这样可以在 URL 中捕获动态值。

```javascript
{
  path: '/users/:userId',
  component: UserProfile
}
```

#### 使用动态参数

在组件中可以使用 `useRoute` 获取动态路由参数：

```html
// UserProfile.vue
<template>
  <div>User ID: {{ userId }}</div>
</template>

<script setup>
import { useRoute } from 'vue-router';

const route = useRoute();
const userId = route.params.userId; // 获取动态路由参数
</script>
```

### 7. 获取和修改查询字符串

#### 1. 获取查询字符串参数

在 Vue 3 组合式 API 中，可以使用 `useRoute` 获取路由信息，包括查询字符串。

```html
<template>
  <div>
    <h1>Query Parameters</h1>
    <p>User ID: {{ userId }}</p>
    <p>Name: {{ name }}</p>
  </div>
</template>

<script setup>
import { useRoute } from 'vue-router';

const route = useRoute(); // 获取路由对象
const userId = route.query.userId; // 获取查询参数
const name = route.query.name; // 获取查询参数
</script>
```

#### 2. 修改查询字符串参数

可以通过路由实例的 `push` 方法来更新查询参数：

```html
<template>
  <button @click="updateQuery">Update User ID</button>
</template>

<script setup>
import { useRouter } from 'vue-router';

const router = useRouter();

const updateQuery = () => {
  router.push({ query: { userId: '456', name: 'Alice' } }); // 更新查询参数
};
</script>
```

### 8. 编程式导航


在 Vue 3 中，可以通过 `useRouter` 来进行编程式导航：

```html
<template>
  <button @click="goToHome">Go to Home</button>
</template>

<script setup>
import { useRouter } from 'vue-router';

const router = useRouter();

const goToHome = () => {
  router.push('/home'); // 导航到主页
};
</script>
```

### 9. 路由守卫:
导航守卫的主要目的是在用户尝试访问需要身份验证的页面时，检查用户是否已登录。如果未登录，则将用户重定向到登录页面；如果已登录，则允许继续访问目标页面。

导航守卫 beforeEach:

`router.beforeEach((to, from, next) => { ... }) ` 是一个全局的导航守卫，它在每次路由切换之前被调用。
- to 是即将要进入的路由对象。
- from 是当前导航正要离开的路由对象。
- next 是一个函数，调用它来继续当前的导航。

```java
// meta 是自定义的, requestAuth 也是自定义的, 用来表示是否需要登录状态才可以访问
{
      path: "/pk",
      name: "pk",
      component: PkIndex,
      meta: {
        requestAuth: true,
      },
    },

router.beforeEach((to, from, next) => {
  // useUserStore 会在下面的 Pinia 中讲到
  const userStore = useUserStore(); // 在beforeEach外面用的话会报错
  if (to.meta.requestAuth && !userStore.user.is_login) {	// userStore.user.is_login表示用户是否已登录
    next({ name: "user_account_login" });
  } else {
    next();
  }
});
```
## 三、组件传递数据

我们在项目中分别创建`Person.vue` 和 `App.vue`, 其中App.vue中使用`<Person/>`创建Person组件实例, 即`App`是`Person`的父组件

在一般情况下我们`父传子`的时候用的是下面讲的第一种`1. 通过defineProps`、`子传父`的时候用的第二种`2. 通过自定义事件`

### 1. 父组件向子组件传递数据

1. 通过`defineProps`

父组件中`<Person :person="person" :str="str" />`， 前面我们向子组件中传递了`person`和`str`变量, 其中`=`右边的`"person"`和`"str"` 是我们在父组件`<script>的setup`中定义的两个变量, `=`左边的是 子组件接受时用到的`name`,  一般命名和父组件中定义时的变量名相同。(个人习惯)

子组件中`const AppProps = defineProps(["person", "str"])`, 其中`[]`中的两个字符串对应父组件中的`name`, AppProps 是包含`person`和`str`的对象， 这样子组件中就可以通过`AppProps.person`访问

### 2. 子组件向父组件传递数据

1. 通过`defineProps`
在上面`父向子`传的时候，可以让父组件给子组件传递一个get方法, 然后让子组件中用`defineProps`接收一下这个方法，在合适的时候调用, 把要传的数据放到get方法的参数列表

2. 通过自定义事件`defineEmits`(不只局限于子组件向父组件, 任意两个组件之间都可以)
父组件中定义了`get-child-name`自定义事件, 给事件绑定了一个函数, 子组件在合适的时机触发`get-child-name`事件即可
```javascript
//父组件中
<Person @get-child-name="getChildName" />
const getChildName = (ChildName) => {
  console.log(ChildName.value)
}

//子组件中
const name = ref('child')
const emit = defineEmits()
const sendName = () => {
  emit('get-child-name', name)
}
```



### 3. mitt自定义事件管理库

上面讲到的`3.2`的第二种自定义事件一般用于子组件和父组件, 但对于层级较深的就需要一层一层传递，很麻烦， mitt就可以很好帮我们解决问题

`mitt`一个`js`的事件管理库, 一般用于管理自定义事件。(安装指令`npm install mitt`)

一般在一个js文件中创建一个mitt实例, 在其他要用到自定义事件管理的文件中引入mitter
```javascript
创建一个新的事件总线实例，返回一个包含 on, off, emit, all 等方法的对象，用于管理事件的订阅和发布。
const mitter = mitt();

订阅指定类型的事件，当该类型的事件被触发时，执行指定的处理函数。
mitter.on('eventName', (data) => {
  console.log('Event received with data:', data);
});

取消订阅指定类型的事件，停止执行特定的处理函数。
mitter.off('eventName', handlerFunction);

触发指定类型的事件，并传递可选的事件数据给订阅该事件的处理函数。
mitter.emit('eventName', { message: 'Hello, mitt!' });

订阅所有类型的事件，当任何类型的事件被触发时执行指定的处理函数。
mitter.all((type, data) => {
  console.log(`Event ${type} received with data:`, data);
});

// 清除所有事件监听器
mitter.all.clear();
```

## 四、组件之间共享数据Pinia

对于子组件和父组件之间，我们可以通过上面讲到的方法实现数据共享，但是对于有更深层级的组件，我们就需要集中管理状态的方法。`Pinia`就可以解决这个问题, 安装命令`npm install pinia`

下面讲到的`pinia`写法更推荐用vue3官方推荐的组合式写法。

### 1. 数据的定义

在实例中引入`pinia`, 跟引入`router`一样
```javascript
import { createPinia } from 'pinia'
const app = createApp(App)
const pinia = createPinia();
app.use(pinia)
```

一般来说我们会把要共享的组件的数据的pinia放到一个`src`文件夹下的`store`中的`js`文件中。然后在需要使用的文件中引入`js`， 生成`store`
```javascript
import { defineStore } from 'pinia'
// defineStore接收的第一个参数是这个store的id(自己定义)
export const usePersonStore = defineStore('Person', () => {
  let name = ref('hh');

  const changeName = (newname) => {
	name = newname;
  }
	
  return {name, changeName};	// 记得把要用到的变量和方法暴漏出去

  // 上面的写法是组合式的, 选项式的话一般把变量卸载state方法的return 中, 而且变量默认就是响应式的, 
  // 组合式的写法, 变量如果要弄成响应式的话需要自己加上 ref 或 reactive 。
})
```
如果要使用`Person`的数据的话就:
```javascript
import { usePersonStore } from '@/store/Person'
const PersonStore = usePersonStore()	//PersonStore 是一个对象
//PersonStore 中包含的就有name, PersonStore.name或PersonStore.$state.name
```
### 2. 数据的修改

1. 直接修改
```javascript
PersonStore.name = 'ld'
```

2. 通过函数修改
```javascript
PersonStore.changeName('ld');	// 
```

对于PersonStore解构出来的数据跟`1.2`中讲的一样也是不再具有响应式，我们可以用`storeToRefs`来解构。(虽然也可以用toRefs，但非常不建议)


4.3 和 4.4 是选项式的写法, 可以仅作了解

### 3. action 和 getters

如果我们需要计算和返回基于已有状态（state）的新数据，那么可以在`getters`中写方法
```javascript
  state() {
    return {
      name: 'yxc',
      age: 18
    }
  },
  action: {
  	// 里面写函数, this指向state
  },
  getters: {	// vue也会自动维护一个this
    doubleName() {
      return this.name + this.name
    }
  }
```

### 4. subscribe

`subscribe` 是用于监听 `store` 的状态变化的方法。Pinia 提供了一种简单而直观的方式来订阅状态变化

订阅方法：通过 `store.$subscribe(callback)` 方法来订阅状态变化，callback 接受一个参数 `mutation`，可以是状态的变化信息或其他相关内容。

取消订阅：订阅方法返回一个取消订阅的函数 `unsubscribe()`，你可以在不需要继续监听状态变化时调用它来取消订阅。

响应式更新：Pinia 使用 Vue.js 的响应式系统来管理状态，因此任何影响状态的变化都会触发订阅者的回调函数。
```javascript
PersonStore.$subscribe((mutation, state) => {})
```
mutation: 一个对象，描述了引起状态变化的 mutation。它包含以下属性：

- storeId: 发生变化的 store 的 ID。
- type: 变化的类型，通常是 direct（直接变化）或 patch（通过 patch 变化）。
- events: 包含变化详情的数组。

state: 变化后的 store 状态。

## 五、组件或Dom元素的引用



在 Vue.js 中，`ref` 属性是一个非常有用的特性，它允许你在组件或 DOM 元素上注册一个引用，以便在 JavaScript 中轻松访问它。这在进行直接操作或与组件/DOM 交互时非常便利。

### 1. `ref` 的基本用法

在 Vue 组件中，你可以在任何元素或子组件上使用 `ref` 属性。例如：

```html
<template>
  <div>
    <input ref="myInput" type="text" />
    <button @click="focusInput">聚焦输入框</button>
  </div>
</template>

<script setup>
import { ref } from 'vue';

const myInput = ref(null); // 创建 ref，初始值为 null

const focusInput = () => {
  myInput.value.focus(); // 使用 .value 访问 DOM 元素
};
</script>
```

在上面的示例中：
- `ref="myInput"` 为输入框注册了一个引用。
- 在 `focusInput` 方法中，通过 `myInput.value` 访问该输入框并调用其 `focus` 方法，使输入框获得焦点。

### 2. 组件引用

如果 `ref` 应用在一个子组件上，`ref` 将会返回该子组件的实例。这样你就可以访问子组件的公开方法和属性。

```html
<template>
  <ChildComponent ref="child" />
  <button @click="callChildMethod">调用子组件方法</button>
</template>

<script setup>
import { ref } from 'vue';
import ChildComponent from './ChildComponent.vue';

const child = ref(null); // 创建 ref，初始值为 null

const callChildMethod = () => {
  child.value.someMethod(); // 使用 .value 访问子组件实例
};
</script>
```

### 3. 使用 `ref` 在 v-for 循环中引用子组件

当 `ref` 应用在 v-for 循环中，Vue 会生成一个包含所有相同 `ref` 名称的数组。

```html
<template>
  <div v-for="item in items" :key="item.id">
    <ChildComponent ref="children" /> <!-- 使用相同的 ref 名称 -->
  </div>
  <button @click="callAllChildrenMethods">调用所有子组件方法</button>
</template>

<script setup>
import { ref } from 'vue';
import ChildComponent from './ChildComponent.vue';

const items = ref([{ id: 1 }, { id: 2 }, { id: 3 }]); // 示例数据
const children = ref([]); // 创建一个 ref 数组

const callAllChildrenMethods = () => {
  children.value.forEach(child => {
    if (child) {
      child.someMethod(); // 调用每个子组件的方法
    }
  });
};
</script>
```



### 4. `ref` 的生命周期

`ref` 的值在组件的生命周期中会随之变化：
- 在组件创建时，`ref` 是空的。
- 在 `mounted` 钩子中，`ref` 已经被填充，可以安全地使用。
- 当组件被销毁时，`ref` 也会被清空。



---

学到这里, 你就可以用vue写项目了。