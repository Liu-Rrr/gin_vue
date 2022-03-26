<template>
<div style="height: 50px; line-height: 50px; border-bottom: 1px solid #ccc; display: flex"><!--盒装弹性布局-->

  <div style="width: 200px;padding-left: 30px; font-weight: bold; color:dodgerblue  ">后台管理</div>
  <div style="flex: 1"></div>
  <div style="width: 100px">
      <el-dropdown>
        <span class="el-dropdown-link">
          {{ nickname }}<i class="el-icon-arrow-down el-icon--right"></i>
        </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item>个人信息</el-dropdown-item>
                <el-dropdown-item @click="$router.push('/login')">退出系统</el-dropdown-item>
              </el-dropdown-menu>
            </template>
      </el-dropdown>
  </div>

</div>


</template>

<script>

export default {
  name: "Header",
  // prop:['user'],
  data(){
    // return{}
    // return
    var str=sessionStorage.getItem("user")
    var json=JSON.parse(str)
    return{
      nickname:json.nickName
    }
  }

}
</script>

<style scoped>

</style>