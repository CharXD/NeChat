<template>
  <div class="Avatar">
    <el-avatar :size="40" :src="friendList.avatar"></el-avatar>
    <p>{{ friendList.username }}</p>
  </div>
</template>

<script>
export default {
  props: ['friendList'],
  name: "FriendAvatar"
}
</script>

<style lang="less" scoped>
.Avatar {
  width: 100%;
  height: 100%;
  border-bottom: 1px solid #ccc;
  margin-top: 10px;

  .el-avatar {
    margin-top: 5%;
    float: left;
    margin-right: 5px;
  }

  p {
    font-size: 15px;
    white-space: nowrap; text-overflow: ellipsis; overflow: hidden;

  }
}
</style>