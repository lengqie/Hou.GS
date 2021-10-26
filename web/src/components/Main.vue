<template>
  <el-container>
    <el-header>
      <h1 class="title">🙈🙉🙊</h1>
    </el-header>
    <el-main>
      <el-form :inline="true">
        <el-form-item>
          <el-input v-model="url" placeholder />
        </el-form-item>
        <el-form-item>
          <el-button @click="getKey" type="primary">提交</el-button>
        </el-form-item>
      </el-form>
    </el-main>
    <el-footer>
      <div>
        <el-alert title="Hou.GS 短网址" type="info" effect="dark" center></el-alert>
        <el-alert
          title="基于Vue + ElementUiPlus + Gin + Gorm + Mysql"
          type="info"
          effect="dark"
          center
        ></el-alert>
        <el-alert title="地址需要 http:// 或 https:// 开头" type="info" effect="dark" center></el-alert>
        <el-alert title="非礼勿视，非礼勿听，非礼勿言" type="info" effect="dark" center></el-alert>
        <el-alert title="2021年10月25日 " type="info" effect="dark" center></el-alert>
      </div>
    </el-footer>
  </el-container>
</template>

<style>
.el-container {
  text-align: center;
  margin: 0 auto;
  margin-top: 10%;
}
.title {
  margin-top: 0;
}
.el-main {
  padding: 0;
}
.el-input {
  width: 500px;
  padding: 0;
  margin-left: 30px;
}
.el-form-item {
  margin: auto;
}
.el-button {
  margin-left: 10px;
  margin-right: 20px;
}
.el-footer {
  width: 610px;
  margin: 0 auto;
}
.el-alert {
  width: 590px;
  margin-top: 10px;
  margin-left: -10px;
}
@media screen and (max-width: 600px) {
  .el-input {
    margin-top: 200px;
    width: 100px;
    margin-left: 20px;
  }
  .title {
    margin-top: 40%;
  }
  .el-alert {
    width: 300px;
    margin-top: 10px;
    margin-left: 20px;
  }
}
</style>

<script>
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
import axios from 'axios'
export default {
  setup() {
    const url = ref("")
    function getKey() {
      if (url.value == "") {
        ElMessage({
          message: '输入框不能为空！',
          type: 'warning',
        })
      } else {
        axios({
          method: "get",
          url: "http://127.0.0.1:8080/v1/url/" + url.value,
        }).then(response => {
          if (response.data.status == 1) {
            ElMessage({
              message: response.data.msg,
              type: 'error',
            })
          } else {
            ElMessage({
              message: response.data.msg,
              type: 'success'
            })
          }
        }).catch(err => {
          console.log(err)
        })
      }
    }
    return {
      url,
      getKey,
    }
  }
}
</script>