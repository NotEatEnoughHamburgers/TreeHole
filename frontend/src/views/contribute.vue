<script setup>
import {getCurrentInstance, h, ref} from "vue";
import {ElNotification} from "element-plus";
import { useRouter} from "vue-router";
const router = useRouter()
// 反馈
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


  let Title = ref()
  let Text = ref()
  const contribute = () => {
    axios.post('/article/create',{
      title:Title.value,
      text:Text.value,
    }).then(res=>{
      msg(res.msg)
      router.push("/")
    }).catch(err=>{
      msg(err.response.data.msg)
    })
  }
</script>

<template>
  <div style="text-align:left;float: left;margin-left: 20%;width: 64%;">
    <h1>我要投稿</h1>
    标题内容：
    <el-input
      v-model="Title"
      placeholder="文章标题"
      style="font-size: 20px;"
    >

    </el-input>
    文本内容：
    <el-input
        v-model="Text"
        :rows="10"
        type="textarea"
        placeholder="投稿内容 后期支持markdown语法"
        style="font-size: 20px;"
    />
    <el-button @click="contribute" style="float:left; margin-top: 20px;background-color: #0371E6;color: white;height: 35px;width: 212px;border: 0px;border-radius: 0.22rem;" type="primary" >投稿</el-button>

  </div>
</template>

<style scoped>

</style>