<template>
  <div class="bg">
    <div class="stage">
      <!--            <el-container>
                    &lt;!&ndash;联系人&ndash;&gt;
                    <el-aside width="200px">
                      <div class="avatar_box">
                        <div class="avatar_top_box">
                          <el-avatar :size="40" src="https://avatars.githubusercontent.com/u/96650292?v=4"></el-avatar>
                          <span>木子日天</span>
                        </div>

                        <div class="main_box">
                          <el-menu :router="true">
                            <el-menu-item :index="'/chat/talk'">
                              <Avatar></Avatar>
                            </el-menu-item>
                            <el-menu-item :index="'/chat/talk'">
                              <Avatar></Avatar>
                            </el-menu-item>
                            <el-menu-item :index="'/chat/talk'">
                              <Avatar></Avatar>
                            </el-menu-item>
                          </el-menu>
                        </div>

                      </div>

                    </el-aside>
                    &lt;!&ndash;聊天&ndash;&gt;
                    <el-main>
                      <router-view></router-view>
                    </el-main>


                  </el-container>-->
      <el-container>
        <!--左侧联系人区域-->
        <el-aside width="200px">
          <user-avatar></user-avatar>
          <el-menu>
            <el-menu-item>
              <friend-avatar ></friend-avatar>
            </el-menu-item>

          </el-menu>
        </el-aside>

        <el-container>
          <!--头部区域-->
          <el-header>Header</el-header>
          <!--主体区域-->
          <el-main>Main</el-main>
          <!--底部发送消息区域-->
          <el-footer>Footer</el-footer>
        </el-container>
      </el-container>
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
  border-radius: 30px;
  overflow: hidden;

  .el-container {
    width: 100%;
    height: 100%;

    .el-header, .el-footer {
      background-color: #B3C0D1;
      color: #333;
      text-align: center;
      line-height: 60px;
    }

    // 左侧联系人区域
    .el-aside {
      background-color: #f7f7f7;
      color: #333;
      text-align: center;
      line-height: 200px;
      border-right: 3px solid #ccc;
    }

    // 主体区域
    .el-main {
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


</style>