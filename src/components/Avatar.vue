<template>
  <div class="Avatar">
    <el-avatar :size="40" src="https://avatars.githubusercontent.com/u/96650292?v=4"></el-avatar>
    <span>木子日天</span>
  </div>
</template>

<script>
export default {
  name: "Avatar"
}
</script>

<style lang="less" scoped>
.Avatar {
  width: 80%;
  height: 50px;
  margin: 20px auto;
  border-bottom: 1px solid #ccc;

  .el-avatar {

  }

  span {
    margin-left: 10px;
    vertical-align: middle;
  }
}
</style>