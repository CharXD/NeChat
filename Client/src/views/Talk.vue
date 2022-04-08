<template>
  <div class="talkBox">
    <ul class="talk_list">
      <li class="left_word">
        <img style="width: 40px;height: 40px;border-radius: 50%"
             src="https://avatars.githubusercontent.com/u/96650292?v=4"/>
        <span>最近还好吗</span>
      </li>
      <li class="right_word">
        <img style="width: 40px;height: 40px;border-radius: 50%"
             src="https://avatars.githubusercontent.com/u/96650292?v=4"/>
        <span>最近还好吗最近还好吗最近还好吗最近还好吗最近还好吗最近还好吗最近还好吗最近还好吗最近还好吗最近还好吗最近还好吗最近还好吗最近还好吗最近还好吗最近还好吗最近还好吗最近还好吗最近还好吗最近还好吗最近还好吗最近还好吗最近还好吗最近还好吗最近还好吗最近还好吗最近还好吗最近还好吗最近还好吗最近还好吗最近还好吗最近还好吗最近还好吗最近还好吗</span>
      </li>
      <li class="right_word">
        <img style="width: 40px;height: 40px;border-radius: 50%"
             src="https://avatars.githubusercontent.com/u/96650292?v=4"/>
        <span>最近还好吗</span>
      </li>
      <li class="right_word">
        <img style="width: 40px;height: 40px;border-radius: 50%"
             src="https://avatars.githubusercontent.com/u/96650292?v=4"/>
        <span>最近还好吗</span>
      </li>
      <li class="right_word">
        <img style="width: 40px;height: 40px;border-radius: 50%"
             src="https://avatars.githubusercontent.com/u/96650292?v=4"/>
        <span>最近还好吗</span>
      </li>
      <li class="right_word">
        <img style="width: 40px;height: 40px;border-radius: 50%"
             src="https://avatars.githubusercontent.com/u/96650292?v=4"/>
        <span>最近还好吗</span>
      </li>
      <li class="right_word">
        <img style="width: 40px;height: 40px;border-radius: 50%"
             src="https://avatars.githubusercontent.com/u/96650292?v=4"/>
        <span>最近还好吗</span>
      </li>
      <li class="right_word">
        <img style="width: 40px;height: 40px;border-radius: 50%"
             src="https://avatars.githubusercontent.com/u/96650292?v=4"/>
        <span>最近还好吗</span>
      </li>

    </ul>
  </div>
</template>

<script>
export default {
  name: "Talk"
}
</script>

<style lang="less" scoped>
.talkBox {
  width: 100%;
  height: 100%;
  background-color: #ccc;

  .talk_list {
    display: flex;
    flex-direction: column;

    background-color: skyblue;

    li {
      overflow: hidden;
      margin: 10px 15px 30px;
    }

    .left_word {
      img {
        float: left;
      }

      span {
        float: left;
        background-color: #5c91fb;
        padding: 10px 15px;
        max-width: 290px;
        border-radius: 0 20px 20px 20px;
        font-size: 16px;
        color: #fff;
        margin-left: 13px;
        position: relative;
        line-height: 24px;
      }
    }

    .right_word {
      img {
        float: right;
      }

      span {
        float: right;
        background-color: #fff;
        padding: 10px 15px;
        max-width: 290px;
        border-radius: 20px 0 20px 20px;
        font-size: 16px;
        color: #000;
        margin-right: 13px;
        position: relative;
        line-height: 24px;
      }
    }

  }
}
</style>