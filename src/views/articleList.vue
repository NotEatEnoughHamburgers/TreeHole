<script setup>
import {defineProps, getCurrentInstance, h, ref} from "vue";
 import {useRouter} from "vue-router";
import {ElNotification} from "element-plus";
 const router = useRouter()
// 引入axios
const {proxy} = getCurrentInstance()
const axios = proxy.$axios

const articleList = ref([])

const commentList = ref([
  {
    Article:'',
    CreatedAt:'',
    ID:'',
    Sender:'',
    SenderInfo:{
      Email:'',
      ID:'',
      ImgUrl:'',
      Name:'',
      Number:'',
      Role:'',
    },
    Text:'',
    Title:'',
  },
])

const msg = (v) => {
  ElNotification({
    title: '消息通知',
    message: h('i', { style: 'color: teal' }, v),
    duration: 2000,
  })
}
  const getHomeComment =() => {
   axios.get('/comment/home').then(res=>{
     commentList.value=res.data
     for (let i = 0; i < commentList.value.length; i++) {
       axios.get('/user/get?id=' + commentList.value[i].Sender).then(res=>{
         commentList.value[i].SenderInfo = res.data
         console.log(res.data)
       })
       for (let j = 0; j < articleList.value.length; j++) {
         if (commentList.value[i].Article === articleList.value[j].ID){
           commentList.value[i].Title =articleList.value[j].Title
         }
       }
     }
   }).catch(err=>{
     console.log(err)
   })

  }
getHomeComment()
 const getArticleList = () => {
   axios.get('/article/list').then(res=>{
     // msg(res.msg)
     // console.log(res.data)
     articleList.value = res.data
   }).catch(err=>{
     msg(err.response.data.msg)
   })
 }
 getArticleList()
 const filterTime = (time)=> {
   return time.substring(0,9)
 }
 const toInfo = (id)=> {
   router.push("info/"+id)
 }
</script>

<template>
  <div style="float: left;margin-left: 3%;width: 64%;">
    <el-card @click="toInfo(item.ID)" v-for="(item,index) in articleList" class="box-card article">
      <h1>{{item.Title}}</h1>
      <div>{{item.Text}}</div>
    </el-card>
    <div style="height: 200px"></div>
  </div>
  <div style="float:right; margin-right: 4%; width: 340px;text-align: left">
    <div class="comment" v-for="(item,index) in commentList">
      <el-avatar
          style="float:left;"
          :src="item.SenderInfo?item.SenderInfo.ImgUrl:'xxx'"
      />
      <div style="">
          {{item.SenderInfo?item.SenderInfo.Name:'未知的用户'}}发表在{{item.Title}}
      </div>
      <div class="time">
        {{filterTime(item.CreatedAt)}}
      </div>
      <div class="font">
        {{item.Text}}
      </div>
    </div>
  </div>
</template>

<style scoped>
.article{
  margin-top: 10px;
  height: 200px;
}
.article:first-child{
  margin-top: 0px;
}
.comment{
  margin-top: 30px;
}

.comment:first-child{
  margin-top: 0px;
}
.comment > div {
  margin-left:  50px;
}
.comment > .time {
  padding-top: 3px;
  font-size: .75em;
}
.comment > .font {
  padding-top: 5px;
  font-size: .875em;
}
</style>