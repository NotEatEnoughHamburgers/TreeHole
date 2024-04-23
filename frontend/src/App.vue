<template>
  <div class="common-layout" v-if="loginStatus">
    <el-container>
      <el-header id="header" class="bg_color" style="position:fixed;right: 0;top:0;left:0;height: 120px; display: flex; align-items: center;">
        <div class="margins">
          <h1 style="float: left;" @click="toHome">山化大树洞</h1>
          <div class="navigation bg_color" style="margin-top: 2%">
            <span class="bg_color" @click="updateLoginStatus">我要登陆</span>
            <span class="bg_color" @click="toContribute">我要投稿</span>
            <span class="bg_color" @click="toHome">首页</span>
          </div>
        </div>
      </el-header>
      <el-main>
        <div style="margin-top: 120px">
          <router-view>
          </router-view>
        </div>
      </el-main>
      <el-footer style="background-color:#F2F5F7;height: 100px; display: flex; align-items: center;position: fixed;left: 0px;bottom: 0px;right: 0px;">
        <div style="background-color:#F2F5F7;margin-left: 6%">
          版权所有 © 2023 - 山东化工职业学院 - 计算机社团
        </div>
      </el-footer>
    </el-container>
  </div>
  <div v-else class="common-layout" style="position: fixed;bottom: 0px;left: 0px;top:0px;right: 0px;">
    <el-container style="height: 100%;">
      <el-aside style="height: 100%; border: 1px solid rebeccapurple" width="40%">
        <h1 style="font-size: 50px;margin-top: 50%">山化大树洞</h1>
      </el-aside>
      <el-main  style="height: 100%; padding: 0">
        <div  style="text-align: left;margin-top: 16%;margin-left: 25%;float:left;width: 100px;">
          <h1 style="margin-left: -5px;font-family:Gabriola,serif;font-size: 35px">登录</h1>
          <div style=" margin-top:-30%;font-size: 20px">
            手机号
          </div>
          <input v-model="number" class="loginInput" />
          <div style="font-size: 20px">
            密码
          </div>
          <input v-model="password" class="loginInput" type="password" />
        </div>
        <div style="margin-top: 50%; float: left;margin-left: -100px">
          <button @click="login" class="loginBtn">登录账号</button>
          <button @click="signupVisibleOpen" class="loginBtn" style="margin-left: 5px">注册账号</button>
        </div>
        <el-dialog
            v-model="signupVisible"
            title="注册账号"
            width="50%"
            :before-close="signupVisibleClose"
        >
          <div style="text-align: left;margin-left: 110px">
            <div style="float:left;">
              <div class="registerText">用户名:</div>
              <br>
              <div class="registerText">手机:</div>
              <br>
              <div class="registerText">密码:</div>
              <br>
              <div class="registerText">验证码:</div>
            </div>
            <div style="float:right;margin-right: 80px">
              <input v-model="username" class="loginInput" />
              <br>
              <input v-model="number" class="loginInput" />
              <br>
              <input v-model="password" class="loginInput" />
              <br>
              <input v-model="code" class="loginInput" />
            </div>

            <div style="float:right;margin-top: 15px;margin-right: 20px">
              <el-button type="primary" @click="register">{{ registerStatus }}</el-button>
              <el-button type="primary" @click="signupVisible=false">前往登录</el-button>
            </div>
            <div style="height: 260px"></div>
          </div>
        </el-dialog>
      </el-main>
    </el-container>
  </div>
</template>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

nav {
  padding: 30px;
}

nav a {
  font-weight: bold;
  color: #2c3e50;
}

nav a.router-link-exact-active {
  color: #42b983;
}

html {
  margin:0;
  padding: 0;
  background-color: #FAFBFC;
}

#header>*{
  background-color: white;
}
.margins{
  margin: 0 6%;width: 100%;
}
.bg_color {
  background-color: white
}
.navigation>span {
  margin-left: 20px;
  float: right;
}

.loginInput{
  /* //margin-left: px; */
  margin-top: 10px;
  border: 2px solid #c3ccd1;
  width: 402px;
  height: 19px;
  border-radius: 0.66rem;
  font-size: 0.95rem;
  padding: 0.6rem 0.75rem;
}
.loginBtn {
  background-color: #0371E6;
  color: white;
  height: 35px;
  width: 212px;
  border: 0px;
  border-radius: 0.66rem;
}
.registerText {
  margin-top: 18px;
}
</style>
<script setup>
import {ref, getCurrentInstance, h} from "vue";
import {ElNotification} from "element-plus";
import { useRouter} from "vue-router";
const router = useRouter()
const msg = (v) => {
  ElNotification({
    title: '消息通知',
    message: h('i', { style: 'color: teal' }, v),
    duration: 2000,
  })
}
// 引入axios
const {proxy} = getCurrentInstance()
const axios = proxy.$axios

const loginStatus = ref(false)
const signupStatus = ref(false)
const signupVisible = ref(false)
const username = ref()
const number = ref()
const password = ref()
const code = ref()
const registerStatus =ref("发送验证码")

const updateLoginStatus =()=> {
  let status = localStorage.getItem('tokenStatus')
  if (status === "success"){
    msg("您已登录")
  } else {
    loginStatus.value = false
  }
}
const signupVisibleOpen =()=>{
  username.value=""
  number.value=""
  password.value=""
  code.value=""
  signupVisible.value = true
}
const signupVisibleClose = () => {
  signupVisible.value = false
  username.value=""
  number.value=""
  password.value=""
  code.value=""
}
const checkLoginStatus = () => {
  let status = localStorage.getItem('tokenStatus')
  let token = localStorage.getItem('jwtToken')
  if (token!==null && status=="success") {
    loginStatus.value = true
  } else {
    if (loginStatus.value === false) {
        return
    }
    msg("登陆状态已失效")
    // ElNotification({
    //   title: '登陆状态',
    //   message: h('i', { style: 'color: teal' }, '登陆状态已失效'),
    //   duration: 2000,
    // })
    loginStatus.value = false
  }
}
checkLoginStatus()

const login = () => {
  axios.post("/user/login", {
    number: number.value,
    pass: password.value,
  }).then(res=>{
    localStorage.setItem('jwtToken',res.data)
    localStorage.setItem('tokenStatus','success')
    msg(res.msg)
    loginStatus.value = true
  }).catch((err)=>{
    msg(err.response.data.msg)
    checkLoginStatus()
  })
}
const register = () => {
  if (registerStatus.value==="发送验证码") {
    // ElNotification({
    //   title: '发送成功',
    //   message: h('i', { style: 'color: teal' }, '发送成功,验证码有效时间为2分钟'),
    //   duration: 2000,
    // })
    registerStatus.value="注册"
    axios.post("/user/sendsms",{
      number:number.value,
    }).then(res=> {
      msg(res.msg)
    }).catch(err => {
      // console.log(err.response)
      msg(err.response.data.msg)
    })
    return
  }
  if (registerStatus.value==="注册") {
    axios.post("/user/signup",{
      name:username.value,
      number:number.value,
      pass:password.value,
      code:code.value,
    }).then(res=>{
      msg(res.msg)
      signupStatus.value= false
    }).catch(err=>{
      msg(err.response.data.msg)
    })
  }
}

// 路由控制
const toContribute = ()=>{
  router.push("contribute")
}
const toHome = ()=>{
  router.push("/")
}

</script>