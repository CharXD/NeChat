<template>
  <div class="bg">
    <div class="stage">

      <!--      <div class="mainBox">
              &lt;!&ndash;左侧联系人区域&ndash;&gt;
              <div class="leftBox">
                <el-aside width="200px">
                  &lt;!&ndash;我的头像&ndash;&gt;
                  <user-avatar></user-avatar>
                  &lt;!&ndash;联系人的头像&ndash;&gt;
                  <div class="friendBox">
                    <el-menu router>
                      <el-menu-item :index="'/chat/talk'" v-for="(i,index) in friendList" :key="index">
                        <Friend-avatar :friendList="i"></Friend-avatar>
                      </el-menu-item>
                    </el-menu>
                  </div>
                  &lt;!&ndash;设置区域&ndash;&gt;

                  <el-button type="primary" icon="el-icon-s-tools">设置</el-button>

                </el-aside>

              </div>
              <el-container>
                &lt;!&ndash;头部区域&ndash;&gt;
                <el-header>
                  Header
                </el-header>
                &lt;!&ndash;主体区域&ndash;&gt;
                <el-main>
                  <router-view></router-view>
                </el-main>
                &lt;!&ndash;底部发送消息区域&ndash;&gt;
                <el-footer>Footer</el-footer>
              </el-container>
            </div>-->

      <!--左边-->
      <div class="leftBox">
        <el-aside width="200px">
          <!--我的头像-->
          <user-avatar></user-avatar>
          <!--联系人头像-->
          <div class="friendBox">
            <el-menu router>
              <el-menu-item :index="'/chat/talk'" v-for="(i,index) in friendList" :key="index">
                <Friend-avatar :friendList="i"></Friend-avatar>
              </el-menu-item>
            </el-menu>
          </div>
          <!--设置区域-->

          <el-button type="primary" icon="el-icon-s-tools">设置</el-button>

        </el-aside>
      </div>
      <!--头部-->
      <div class="headerBox">
        木子日天
      </div>
      <!--      中间-->
      <div class="mainBox">
        <router-view></router-view>
      </div>
      <!--      底部-->
      <div class="footerBox">

      </div>
    </div>

  </div>
</template>

<script>
import FriendAvatar from "../components/FriendAvatar";
import UserAvatar from "@/components/UserAvatar";

export default {
  name: "chat.vue",
  data() {
    return {
      friendList: [
        {id: 1, username: '木子日天木子日天', avatar: 'https://avatars.githubusercontent.com/u/96650292?v=4'},
        {id: 2, username: '洋洋洋洋洋洋洋洋', avatar: 'https://avatars.githubusercontent.com/u/28512134?v=4'},
        {id: 3, username: '充电器', avatar: 'https://avatars.githubusercontent.com/u/39784843?v=4'},
        {id: 3, username: '充电器', avatar: 'https://avatars.githubusercontent.com/u/39784843?v=4'},
        {id: 3, username: '充电器', avatar: 'https://avatars.githubusercontent.com/u/39784843?v=4'},
        {id: 3, username: '充电器', avatar: 'https://avatars.githubusercontent.com/u/39784843?v=4'},
        {id: 3, username: '充电器', avatar: 'https://avatars.githubusercontent.com/u/39784843?v=4'},
        {id: 3, username: '充电器', avatar: 'https://avatars.githubusercontent.com/u/39784843?v=4'},
        {id: 3, username: '充电器', avatar: 'https://avatars.githubusercontent.com/u/39784843?v=4'},
        {id: 1, username: '木子日天木子日天', avatar: 'https://avatars.githubusercontent.com/u/96650292?v=4'},
        {id: 1, username: '木子日天木子日天', avatar: 'https://avatars.githubusercontent.com/u/96650292?v=4'},
        {id: 1, username: '木子日天木子日天', avatar: 'https://avatars.githubusercontent.com/u/96650292?v=4'},

      ],
      sendMsg: {
        message: "",

      }
    }
  },
  components: {FriendAvatar, UserAvatar},
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
  position: relative;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  height: 650px;
  width: 1100px;
  background-color: #f7f7f7;
  border-radius: 20px;
  overflow: hidden;

  .el-container {
    position: relative;
    width: 100%;
    height: 100%;

    .el-header, .el-footer {
      position: absolute;
      width: 900px;
      right: 0;
      background-color: #B3C0D1;
      color: #333;
      text-align: center;
      line-height: 60px;
    }


    // 主体区域
    .el-main {
      position: absolute;

      background-color: #E9EEF3;
      color: #333;
      text-align: center;
      line-height: 160px;
    }

    body > .el-container {
      margin-bottom: 40px;
    }

    .el-container:nth-child(5) .el-aside,
    .el-container:nth-child(6) .el-aside {
      line-height: 260px;
    }

    .el-container:nth-child(7) .el-aside {
      line-height: 320px;
    }
  }
}

.leftBox {
  position: absolute;
  left: 0;
  background-color: pink;
  width: 200px;
  height: 100%;
  // 左侧联系人区域
  .el-aside {
    height: 100%;
    width: 200px;
    position: absolute;
    background-color: #f7f7f7;
    color: #333;
    text-align: center;
    border-right: 3px solid #ccc;
    overflow-y: hidden;
    // 联系人
    .friendBox {
      width: 100%;
      height: 80%;
      overflow-y: scroll;
      overflow-x: hidden;
      margin-bottom: 0;
    }

    .el-menu {
      width: 100%;
      height: 100%;
    }

    .el-button {
      position: absolute;
      left: 0;
      bottom: 0;
      width: 100%;
      height: 50px;
      font-size: 16px;
    }

  }
}

.headerBox {
  text-align: center;
  line-height: 30px;
  font-size: 20px;
  position: absolute;
  width: 900px;
  height: 30px;
  background-color: #ccc;
  right: 0;
}

.mainBox {
  position: absolute;
  top: 30px;
  right: 0;
  width: 900px;
  height: 500px;
  background-color: skyblue;
  overflow-y: scroll;
}

.footerBox {
  position: absolute;
  width: 900px;
  height: 120px;
  background-color: purple;
  bottom: 0;
  right: 0;
}

</style>