<script setup>
import {defineProps, getCurrentInstance, h, ref} from "vue";
import {ElNotification} from "element-plus";
  // 引入axios
  const {proxy} = getCurrentInstance()
  const axios = proxy.$axios

  let articleData = ref({
    CreatedAt:'',
    DeletedAt:'',
    ID:'',
    Modified:'',
    Sender:'',
    Text:'',
    Title:'',
    UpdatedAt:'',
  })
  let user = ref({
    Email:'',
    ID:'',
    ImgUrl:'',
    Name:'',
    Number:'',
    Role:''
  })
  const msg = (v) => {
    ElNotification({
      title: '消息通知',
      message: h('i', { style: 'color: teal' }, v),
      duration: 2000,
    })
  }

let commentList = ref([{
  Article:'',
  CreatedAt:'',
  DeletedAt:'',
  ID:'',
  Parent:'',
  Sender:'',
  Text:'',
  SenderInfo:{
    ID:'',
    Name:'',
    Number:'',
    ImgUrl:'',
    Email:'',
    Role:'',
  },
  subset:[{
    Article:'',
    CreatedAt:'',
    DeletedAt:'',
    ID:'',
    Parent:'',
    Sender:'',
    Text:'',
    SenderInfo:{
      ID:'',
      Name:'',
      Number:'',
      ImgUrl:'',
      Email:'',
      Role:'',},
  },],
  reply:false,
}])
  let props = defineProps({
    id:Number,
  })
  const getUser = (id) => {
    axios.get('/user/get?id='+id).then(res=>{
      user.value= res.data
    }).catch(err=>{
      msg(err.response.data.msg)
    })
  }
  const getArticle = ()=> {
    axios.get('/article/get?id='+props.id).then(res=>{
      articleData.value=res.data
      user.value=(articleData.value.Sender)
      getUser(articleData.value.Sender)
    }).catch(err=>{
      msg(err.response.data.msg)
    })
  }
  getArticle()
const getComment = async () => {
  try {
    const res = await axios.get('/comment/list?article=' + props.id);
    const commentListData = res.data;

    // 使用 Promise.all 来并行请求用户详细信息
    await Promise.all(commentListData.map(async (comment) => {
      const userRes = await axios.get('/user/get?id=' + comment.Sender);
      comment.SenderInfo = userRes.data;
    }));

    // 组织子评论
    const organizedComments = [];
    for (let i = 0; i < commentListData.length; i++) {
      const subset = commentListData.filter(comment => comment.Parent === commentListData[i].ID);
      commentListData[i].subset = subset;
      if (!commentListData[i].Parent) {
        organizedComments.push(commentListData[i]);
      }
    }

    // 更新 commentList.value
    commentList.value = organizedComments;

  } catch (error) {
    console.error('获取评论失败:', error);
    // 处理错误，例如显示错误信息
    // msg(error.response.data);
  }
};
  getComment()
  const formattedText = (value) => {
    return value.replace(/\n/g, '<br>');
  }
  const formattedTime = (time)=> {
    return time.substring(0,9)
  }
  const formattedTimeHour = (time)=> {
    return time.substring(11,19)
  }
  const getLength = (v)=> {
    return v.length
  }

  const replyStatus = ref()
  // 切换回复部分的显示/隐藏状态
  const toggleReply = (index) => {
    for (let i = 0; i < commentList.value.length; i++) {
      commentList.value[i].reply = false
    }
    if (replyStatus.value===index){
      commentList.value[index].reply = false
      replyStatus.value = undefined
    } else {
      commentList.value[index].reply = true
      replyStatus.value = index
    }
  };
  const replyText =ref()
  const submitReply = (id) => {
    let response
    if (id) {
      console.log(1);
      response = axios.post("/comment/send",{
        article:props.id,
        text:replyText.value,
        parent:id,
      })
    }else{
      response = axios.post("/comment/send",{
        article:props.id,
        text:replyText.value,
      })
    }
    response.then(res=>{
      msg(res.msg)
      replyText.value = ""
      getComment()
    }).catch(err=>{
      msg(err.response.data.msg)
    })
  }
</script>

<template>
  <div style="margin: 0 23%">
    <h1 style="font-size: 30px" class="left">{{articleData.Title}}</h1>
    <el-avatar
        style="float:left;height: 25px;width: 25px;"
        :src="user.ImgUrl"
    />
    <el-breadcrumb style="margin-top: 30px;margin-left: 35px" separator="/">
      <el-breadcrumb-item>山化大树洞</el-breadcrumb-item>
      <el-breadcrumb-item>{{ formattedTime(articleData.CreatedAt) }}</el-breadcrumb-item>
    </el-breadcrumb>
    <div style="color: #FF9541">『感谢<b style="font-weight:bold;color: #FF8C00">{{user.Name}}</b>同学的投递』</div>
    <div style="margin-top: 20px" class="left" v-html="formattedText(articleData.Text)"></div>

    <div style="margin-top: 100px;" class="left">
      <h1 style="font-size: 30px;">{{getLength(commentList)}} 评论</h1>
<!--      一级评论-->
      <div v-for="(item,index) in commentList">
        <hr style="border: none;border-top: 1px solid #ccc;margin: 50px 0;">
        <el-avatar
            v-if="item.SenderInfo && item.SenderInfo.ImgUrl"
            style="float: left;height: 50px;width: 50px;"
            :src="item.SenderInfo.ImgUrl"
        />
        <el-button class="replyBtn" @click="toggleReply(index)" style="float: right" size="small" type="primary" plain>
          {{ item.reply ? '收起回复' : '回复' }}
        </el-button>
        <div style="float: right;margin-right: 5px;margin-top: 5px">
          <el-breadcrumb  separator="/">
            <el-breadcrumb-item>{{ formattedTime(item.CreatedAt) }}</el-breadcrumb-item>
            <el-breadcrumb-item>{{ formattedTimeHour(item.CreatedAt) }}</el-breadcrumb-item>
          </el-breadcrumb >
        </div>
        <div>
          <div class="commentFont" style="font-size: 16px;font-weight: 700;">
            {{ item.SenderInfo ? item.SenderInfo.Name : '未知用户' }}
          </div>
          <div class="commentFont">{{item.Text}}</div>
        </div>
<!--        二级-->
        <div style="margin-left: 50px" v-for="(item2,index2) in item.subset">
          <hr style="border: none;border-top: 1px solid #ccc;margin: 50px 0;">
          <el-avatar
              v-if="item2.SenderInfo && item2.SenderInfo.ImgUrl"
              style="float: left;height: 50px;width: 50px;"
              :src="item2.SenderInfo.ImgUrl"
          />
          <div style="float: right;margin-right: 5px;margin-top: 5px">
            <el-breadcrumb  separator="/">
              <el-breadcrumb-item>{{ formattedTime(item2.CreatedAt) }}</el-breadcrumb-item>
              <el-breadcrumb-item>{{ formattedTimeHour(item2.CreatedAt) }}</el-breadcrumb-item>
            </el-breadcrumb >
          </div>
          <div>
            <div class="commentFont" style="font-size: 16px;font-weight: 700;">
              {{ item2.SenderInfo ? item2.SenderInfo.Name : '未知用户' }}
            </div>
            <div class="commentFont">{{item2.Text}}</div>
          </div>
        </div>
        <!-- 回复部分 -->
        <div v-if="item.reply" style="margin-left: 60px">
          <hr style="border: none;border-top: 1px solid #ccc;margin: 50px 0;">
          <h1>回复给{{item.name}}</h1>
          <el-input
              v-model="replyText"
              :rows="5"
              type="textarea"
              placeholder="输入你的回复"
              style="font-size: 20px"
          />
          <el-button @click="submitReply(item.ID)" style="margin-top: 20px;background-color: #0371E6;color: white;height: 35px;width: 212px;border: 0px;border-radius: 0.22rem;" type="primary" >发表评论</el-button>
        </div>
      </div>
    </div>
    <div>
      <hr style="border: none;border-top: 1px solid #ccc;margin: 50px 0;">
      <h1>留下评论</h1>
      <el-input
          v-model="replyText"
          :rows="5"
          type="textarea"
          placeholder="输入你的回复"
          style="font-size: 20px"
      />
      <el-button @click="submitReply(undefined)" style="float:left; margin-top: 20px;background-color: #0371E6;color: white;height: 35px;width: 212px;border: 0px;border-radius: 0.22rem;" type="primary" >发表评论</el-button>
    </div>
  </div>
  <div style="height: 200px"></div>
</template>

<style scoped>
.left {
  text-align: left;
}
.commentFont{
  margin-top: 15px;
  margin-left: 60px;
}
.commentFont:last-child{
  margin-top: 20px;
}
.replyBtn:active,.replyBtn:focus{
  background-color: #ECF5FF;
  color: #409EFF;
  border-color: #A0CFFF;
}
.replyBtn--primary:hover{
  background-color: #ECF5FF;
  color: #409EFF;
}

</style>