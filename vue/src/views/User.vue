
<template>

  <div style="padding:5px">

    <div style="margin: 5px 0">
      <el-button type="primary" @click="add()">新增</el-button>
      <el-button type="primary">导入</el-button>
      <el-button type="primary">导出</el-button>
    </div>

    <div style="margin: 10px 0">
      <el-input v-model="search" placeholder="请输入关键字" style="width: 20%" ></el-input>
      <el-button type="primary" style="margin-left:2px" @click="load">查询</el-button>
    </div>


    <el-table :data="tableData"
              border style="width: 100%"
              stripe="true">
      <el-table-column prop="id"
                       label="ID"

                       sortable> </el-table-column>
      <el-table-column prop="username"
                       label="姓名"> </el-table-column>

      <el-table-column prop="age"
                       label="年龄"> </el-table-column>

      <el-table-column prop="sex"
                       label="性别"> </el-table-column>

      <el-table-column prop="address"
                       label="地址"> </el-table-column>


      <el-table-column label="操作">
        <template #default="scope">
          <el-button @click="handleEdit(scope.row)" size="mini">编辑
          </el-button>
          <el-popconfirm title="确定删除吗？" @confirm="handleDelete(scope.row.id)">
            <template #reference>
              <el-button  type="warning" size="mini">编辑
              </el-button>
            </template>
          </el-popconfirm>
        </template>
      </el-table-column>

    </el-table>
    <div>
      <el-pagination
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
          :current-page="currentPage"
          :page-sizes="[5, 10, 20]"
          :page-size="pageSize"
          layout="total, sizes, prev, pager, next, jumper"
          :total="total"
      >
      </el-pagination>

      <el-dialog title="提示"  v-model="dialogVisible" width="30%">
        <el-form :model="form" label-width="120px">
          <el-form-item label="姓名" style="width: 80%">
            <el-input v-model="form.username"></el-input>
          </el-form-item>
          <el-form-item label="昵称" style="width: 80%">
            <el-input v-model="form.nickName"></el-input>
          </el-form-item>
          <el-form-item label="年龄"  style="width: 80%">
            <el-input v-model="form.age"></el-input>
          </el-form-item>
          <el-form-item label="性别" style="width: 80%">
            <el-radio v-model="form.sex" label="男">男</el-radio>
            <el-radio v-model="form.sex" label="女">女</el-radio>
            <el-radio v-model="form.sex" label="未知">未知</el-radio>
          </el-form-item>
          <el-form-item label="地址" style="width: 80%">
            <el-input v-model="form.address" ></el-input>
          </el-form-item>
        </el-form>
        <template #footer>
    <span class="dialog-footer">
      <el-button @click="dialogVisible = false">取 消</el-button>
      <el-button type="primary" @click="save">确 定</el-button>
    </span>
        </template>
      </el-dialog>

    </div>
  </div>

</template>

<script>


import request from "@/untils/request";

export default {
  name: 'Home',

  components: {

  },

  data(){
    return{
      //from表单一定要定义，要不然没有提交的弹窗
      form:{},
      dialogVisible:false,
      currentPage:1,
      pageSize:10,
      total:0,
      search:'',
      tableData:[]

    }
  },

  //创造时调用的方法
  created() {
    //调用load()
    this.load()
  },

  methods: {
    //将form保存到后台，前端和后端的数据交互

    load() {
      //config部分要先打{}，否则默认config没有，必须要先填{}
      request.get("/user/count",{

        params:{
          pageNum:this.currentPage,
          pageSize:this.pageSize,
          search:this.search
        }

      }).then(res=>{
        console.log(res)
        this.tableData=res.data.records
        this.total=res.data.total
      })
    },

    add(){
      this.dialogVisible=true
      this.form={}
      //必须为空，要不然之前打的信息还在
    },

    save(){
      if(this.form.id){//更新

        request.put("/user",this.form).then(res=>{
          console.log(res)
          //element的弹窗,res为后台返回的数据，没有data
          if(res.code === '0'){
            this.$message({
              type:"success",
              message:"更新成功"
            })
            this.load()
          }else {
            this.$message({
              type:"failed",
              message:"更新失败"
            })
          }
        })

      }else{//新增

        request.post("/user/add",this.form).then(res=>{
          console.log(res)
          if(res.code === '0'){
            this.$message({
              type:"success",
              message:"新增成功"
            })
            this.load()
          }else{
            this.$message({
              type:"failed",
              message:"新增失败"
            })
          }
        })

      }


      this.dialogVisible=false
    },

    //前端绑定的对象，选中这一行的对象
    handleEdit(row) {
      // console.log(row)
      //深拷贝，避免直接赋值给表单，改变了值按取消后还是改变这一行的值
      this.form=JSON.parse(JSON.stringify(row))
      this.dialogVisible=true

    },

    //html绑定传来的id
    handleDelete(id) {
      //删除id，将id传入后端，后端写个接口删除
      console.log(id)
      //通过这种path传入后端的delete映射中
      request.delete("/user/"+id).then(res=>{

        if(res.code === '0'){
          this.$message({
            type:"success",
            message:"删除成功"
          })
          this.load()
        }else{
          this.$message({
            type:"failed",
            message:"删除失败"
          })
        }
      })

    },

    handleSizeChange(val) {
      //改变每页个数触发
      console.log(`每页 ${val} 条`)
      this.pageSize=val
      this.load()
    },

    handleCurrentChange(val) {
      //改变每页数触发
      this.currentPage=val
      this.load()
      console.log(`当前页: ${val}`)
    },

  },


}

</script>
