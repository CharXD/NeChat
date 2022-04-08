<template>
  <div class="userAvatar">
    <el-avatar :size="50" src="https://avatars.githubusercontent.com/u/96650292?v=4"></el-avatar>
    <span>木子日天</span>
  </div>
</template>

<script>
export default {

  name: "UserAvatar"
}
</script>

<style scoped lang="less">
.userAvatar {
  width: 100%;
  height: 70px;
  background-color: #eeebdd;
  line-height: 80px;
  border-bottom: 1px solid #ccc;

  .el-avatar {
    margin-left: -20px;
  }

  span {
    font-size: 18px;
    margin-left: 10px;
    vertical-align: middle;
  }
}
</style>