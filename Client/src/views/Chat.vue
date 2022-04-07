<template>
  <div class="bg">
    <div class="stage">
      <el-container>
        <!--联系人-->
        <el-aside width="200px">
          <div class="avatar_box">
            <div class="avatar_top_box">
              <el-avatar :size="40" src="https://avatars.githubusercontent.com/u/96650292?v=4"></el-avatar>
              <span>木子日天</span>
            </div>

            <div class="main_box">
              <el-menu :router="true">
                <el-menu-item :index="'/talk'">
                  <Avatar></Avatar>
                </el-menu-item>
                <el-menu-item :index="'/talk'">
                  <Avatar></Avatar>
                </el-menu-item>
                <el-menu-item :index="'/talk'">
                  <Avatar></Avatar>
                </el-menu-item>
              </el-menu>
            </div>
            <!--底部设置按钮-->
            <div class="footer_box">

              <el-button type="primary" @click="goSetting">
                <i class="el-icon-s-tools"></i> 设置
              </el-button>
            </div>
          </div>

        </el-aside>
        <!--聊天-->
        <el-main>
          <router-view></router-view>
        </el-main>


      </el-container>
    </div>

  </div>
</template>

<script>
import Avatar from "../components/Avatar.vue";

export default {
  name: "chat.vue",
  data() {
    return {
      sendMsg: {
        message: "",
      }
    }
  },
  components: {Avatar,},
  methods: {
    goSetting() {
      this.$router.push('/setting')
    }
  }
}
</script>

<style scoped lang="less">
.bg {
  position: absolute;
  top: 0;
  left: 0;
  height: 100%;
  width: 100%;
  background: url(../assets/img/background_chat.jpg) no-repeat fixed center center;
  background-size: cover;
}

.stage {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  height: 650px;
  min-width: 1100px;
  //background-color: #d0c6bd;
  //box-shadow: 0 0 10px 5px #cac0b7;
  border-radius: 30px;
}

.el-container {
  height: 100%;
  width: 100%;
}

.el-aside {
  height: 100%;
  background-color: #fff;
  overflow: hidden;
}

.el-main {
  height: 100%;
  background: pink;
  width: 100%;
  border-left: #ccc solid 3px;
  padding: 0;
  overflow-y: auto;
}



// 左侧头像盒子
.avatar_box {
  width: 100%;

}

.el-menu-item {
  width: 100%;
  font-size: 16px;
  padding: 0 !important;
  line-height: 45px;
}

// 自己的头像和名字
.avatar_top_box {
  width: 100%;
  height: 70px;
  background-color: #EEEBDD;
  border-bottom: solid 1px #ccc;

  .el-avatar {
    margin-left: 10px;
  }

  span {
    line-height: 70px;
    text-align: center;
    vertical-align: middle;
    margin-left: 15px;
    font-size: 16px;
  }
}

.el-menu-item {
  width: 100%;
 margin-top: 8px;
}

// 联系人
.main_box {
  height: 540px;
  overflow-y: scroll;
  background-color: #F7F7F7;

}

// 底部设置
.footer_box {
  .el-button {
    border-radius: 0;
    width: 100%;
    font-size: 15px;
  }
}


</style>