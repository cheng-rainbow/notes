Electron 是一个由 GitHub 开发的开源框架，用于构建跨平台的桌面应用程序。它结合了 [Chromium](https://x.com/i/grok?text=Chromium) 和 Node.js 的能力，使得开发者可以使用 JavaScript、HTML 和 CSS 这些前端技术来创建桌面应用。



## 一、基本内容

- **主进程 (Main Process)**：运行在 Node.js 环境中，负责创建浏览器窗口、管理应用生命周期、与操作系统交互等。这个进程中你可以使用 Node.js 的所有模块。

  - main.js 或 index.js 通常是主进程的入口文件。由 `package.json` 中的 main 属性决定。

- **渲染进程 (Renderer Process)**：渲染进程可以访问 DOM 和所有的 Web API，但不能直接访问 Node.js 的 API。这些进程是沙盒化的，主要负责 UI 的渲染。如果需要访问 nodejs 下的api，则需要跟主进程通信，由主进程完成。

  

### 1. 准备工作

#### 1.1 下载electron

```bash
// 1. 初始化一个 nodejs 项目
npm init -y	

// 2. 设置镜像
npm config set registry https://registry.npmmirror.com
$env:ELECTRON_MIRROR="https://npmmirror.com/mirrors/electron/"

// 3. 以依赖的方式下载electron
npm install --save-dev electron	

// 4. 运行项目
npx electron . 

// 4. 运行项目，也可以在package.json的script中加上 "start": "electron ."，之后就可以用下面这个了
npm start

// 4. 以实时查看的方式，运行项目
npm install --save-dev electronmon
// package.json 中
"start": "electronmon ."
```



#### 1.2 主要事件

Electron 应用程序生命周期中的关键点，每个事件都有其特定的用途和触发时机。以下是这些事件的详细解释：

**ready**

- **触发**：当 Electron 完成初始化并准备好创建浏览器窗口时触发。
- **用途**：这是创建窗口和设置应用配置的最佳时机。你应该在这里或 app.whenReady() 之后创建你的主窗口。

**dom-ready**

- **触发**：在渲染进程中，当 DOM 已经准备好（即 DOMContentLoaded 事件发生）时触发。
- **用途**：这是一个渲染进程事件（不是应用级事件），可以通过 webContents 对象监听，用于在 HTML 文档结构加载后但在所有资源（如图片、样式表）加载前执行操作。

**did-finish-load**

- **触发**：当页面加载完成（包括所有资源）时触发。
- **用途**：这也是一个渲染进程事件，适用于在页面完全加载后执行操作，例如注入脚本、修改 DOM 等。

**before-quit**

- **触发**：在应用准备退出之前触发，通常在用户尝试通过菜单或快捷键退出应用时。
- **用途**：可以在这里执行一些清理任务，或者询问用户是否真的要退出。如果你不想要应用退出，可以在这阻止。

**will-quit**

- **触发**：当应用即将退出时触发。这意味着所有窗口都已关闭，或者用户明确表示要退出（如在 macOS 上使用 Cmd + Q）。
- **用途**：这是最后一次可以在应用退出前执行某些操作的机会，比如保存状态、日志记录等。

**quit**

- **触发**：应用完全退出后触发。
- **用途**：这是一个通知事件，表示应用已经退出，你不能再在这做任何操作，这通常用于日志记录或通知其他进程。

**window-all-closed**

- **触发**：当所有浏览器窗口都已关闭时触发。
- **用途**：在非 macOS 系统上，这里通常调用 app.quit() 来退出应用，因为在这些平台上，关闭最后一个窗口通常表示用户希望退出应用。然而，在 macOS 上，用户可能期望通过 Cmd + Q 退出，所以你可能不在这里退出应用。



### 2. 样例代码

下面给一个样例，包含创建一个应用窗口（主进程），在应用窗口上渲染 **html**，通过 **html** 的按钮创建另一个窗口（渲染进程），并显示内容。这个样例包含了基本的内容和 主进程 和 渲染进程之间的通信。



1. main.js文件 （主进程代码）

```javascript
const { app, BrowserWindow, ipcMain } = require("electron");
const path = require('path');

// 创建一个窗口，并配置一些基本内容
function createWindow() {
    const mainWind = new BrowserWindow({
        x: 100,
        y: 100,
        show: false,    // 创建后不立即显示窗口，手动调用mainWind.show()后显示
        width: 1280,
        height: 720,
        maxHeight: 800,
        maxWidth: 1280,
        minHeight: 720,
        minWidth: 1280,
        icon: "./public/icon.ico",  // 应用图标
        title: "test",  // 应用title
        resizable: true,    // 是否可调整窗口大小
        fullscreen: false,  // 是否全屏
        autoHideMenuBar: true,  // 是否自动隐藏菜单栏
        webPreferences: {
            // 这两个注释取消掉后，渲染进程可直接使用electron
            // nodeIntegration: true, // 启用 Node.js 集成
            // contextIsolation: false, // 禁用上下文隔离（与 nodeIntegration: true 配合使用）
            preload: path.join(__dirname, 'preload.js')	// 指定预加载脚本的路径
        }
    });
    mainWind.loadFile("index.html");	// 当前窗口显示什么内容

    mainWind.on("ready-to-show", () => {	// 当窗口加载完成后执行，手动调用show()显示前面创建的窗口
        mainWind.show();
    });
}

// 当 app 准备好后执行
app.whenReady().then(() => {
    createWindow();	// 调用上面函数
	
    // 给ipc绑定事件，当渲染进程通过 IPC 发送一个 "create-new-window" 事件时，主进程会响应这个事件
    ipcMain.on("create-new-window", () => {
        const newWind = new BrowserWindow({
            width: 120,
            height: 100,
            webPreferences: {
                contextIsolation: true,
                preload: path.join(__dirname, 'preload.js'),
            }
        });
        newWind.loadFile("index.html");
    });
});
```

2. preload.js （预加载脚本）

```javascript
// 这个文件实际上运行在主进程中，因为上面我们设置了 ”preload: path.join(__dirname, 'preload.js'),“ 
const { contextBridge, ipcRenderer } = require("electron");

// 通过 contextBridge 将一个 API 暴露到渲染进程的全局 window 对象上。
contextBridge.exposeInMainWorld("testApi", {
    createNewWindow : () => {
        ipcRenderer.send("create-new-window");	// 通过 ipcRender 发送 create-new-window 事件
    },
})
```

3. index.html（窗口显示的内容）

```javascript
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title></title>
</head>
<body>
    <h1>你好!~</h1>
    <button id="btn">点击创建窗口</button>
    <script src="./index.js"></script>
</body>
</html>
```

4. index.js （html的js文件）

```javascript
window.addEventListener("DOMContentLoaded", () => {
    const btn = document.getElementById("btn");
    btn.addEventListener("click", () => {	// 给 html 中的button绑定事件
        window.testApi.createNewWindow();	// 调用 preload.js 文件中的 createNewWindow
    });
})
```



### 3. 自定义菜单

```javascript
const { app, BrowserWindow, Menu} = require("electron");


function createWindow() {
    const mainWind = new BrowserWindow({
        show: false,    // 创建时隐藏窗口，手动调用mainWind.show()显示
        width: 1280,
        height: 720,
    });
    mainWind.loadFile("index.html");

    mainWind.on("ready-to-show", () => {
        mainWind.show();
    });

    const template = [
        {
            label: '文件',  // 菜单栏名
            submenu: [      // 菜单栏下的子菜单
                // role 属性用于指定菜单项的预定义行为。
                // 这些预定义的角色允许 Electron 自动处理某些常见操作的逻辑，
                // 而不需要开发者手动编写代码来实现这些功能。
                { role: 'quit' }
            ]
        },
        {
            label: '编辑',
            submenu: [
                { role: 'undo' },
                { role: 'redo' },
                { type: 'separator' },
                { label : '剪切', role: 'cut' },
            ]
        },
        {
            label: '视图',
            submenu: [
                { role: 'reload' },
                { role: 'forcereload' },
                { role: 'toggledevtools' },
                // 分隔符
                { type: 'separator' },
                { role: 'resetzoom' },
            ]
        },
        {
            label: '窗口',
            submenu: [
                { role: 'minimize' },
                { role: 'close' }
            ]
        },
        {
            label: '帮助',
            role: 'help',
            submenu: [
                {
                    label: '了解更多',
                    // 被点击时触发
                    click () { require('electron').shell.openExternal('https://electronjs.org') }
                }
            ]
        }
    ];
    const menu = Menu.buildFromTemplate(template);
    mainWind.setMenu(menu);
}

app.whenReady().then(() => {
    createWindow();
});
```



### 4. 自定义dialog

```javascript
// 1. 主进程中
ipcMain.on("show-dialog", () => {
        dialog.showMessageBox(mainWind, {
            type: "info",
            title: "提示",
            message: "这是一个提示框",
            buttons: ["关闭"]
        });
        dialog.showOpenDialog(mainWind, {
            defaultPath: "C:\\",
            properties: ["openDirectory", "multiSelections"],
            title : "选择文件"
        }
        );
    })

// 2. preload.js中
const { contextBridge, ipcRenderer, dialog } = require("electron");

contextBridge.exposeInMainWorld("testApi", {
    showDialog : () => {
        ipcRenderer.send("show-dialog"); 
    }
})

// 3. index.js中
window.addEventListener("DOMContentLoaded", () => {
    const btn = document.getElementById("btn");
    btn.addEventListener("click", () => {	
        window.testApi.showDialog();
    });
})
```



### 5. 快捷键绑定

```javascript
app.on("ready", () => {
    globalShortcut.register('ctrl + q', () => {
        console.log('ctrl + q is pressed')
    })
})

app.on("will-quit", () => {
    globalShortcut.unregisterAll();
})
```

