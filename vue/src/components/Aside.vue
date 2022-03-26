<template>
  <div>
    <el-menu
        class="el-menu-vertical-demo"
        default-active="user"
        router
        default-openeds="[1]"
        style="width: 150px; min-height: calc(100vh - 50px)"
    >
      <!-- router为路由跳转-->
<!--      100vh为整个视图高度，-50px是因为减去头部高度， 必须要用空格隔开-号-->
      <el-sub-menu index="1">
        <template #title>
          <i class="el-icon-location"></i>
          <span>系统管理</span>
        </template>
          <el-menu-item index="user">用户管理</el-menu-item>
      </el-sub-menu>
          <el-menu-item index="book">数据管理</el-menu-item>
    </el-menu>
  </div>
</template>

<script>
export default {
  name: "Aside"
}
</script>

<style scoped>

</style>