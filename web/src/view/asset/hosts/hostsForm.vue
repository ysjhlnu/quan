<template>
<div>
    <el-form :model="formData" label-position="right" label-width="80px">
             <el-form-item label="实例ID:">
                <el-input v-model="formData.instanceid" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="实例名字:">
                <el-input v-model="formData.name" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="内网IP:">
                <el-input v-model="formData.privateip" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="地区:">
                <el-input v-model="formData.region" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="环境:">
                <el-input v-model="formData.env" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="运行状态:">
                <el-input v-model="formData.status" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="实例模板:">
                <el-input v-model="formData.type" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="其他:">
                <el-input v-model="formData.ps" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           <el-form-item>
           <el-button @click="save" type="primary">保存</el-button>
           <el-button @click="back" type="primary">返回</el-button>
           </el-form-item>
    </el-form>
</div>
</template>

<script>
import {
    createHosts,
    updateHosts,
    findHosts
} from "@/api/hosts";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "Hosts",
  mixins: [infoList],
  data() {
    return {
      type: "",formData: {
            instanceid:"",
            name:"",
            privateip:"",
            region:"",
            env:"",
            status:"",
            type:"",
            ps:"",
            
      }
    };
  },
  methods: {
    async save() {
      let res;
      switch (this.type) {
        case "create":
          res = await createHosts(this.formData);
          break;
        case "update":
          res = await updateHosts(this.formData);
          break;
        default:
          res = await createHosts(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"创建/更改成功"
        })
      }
    },
    back(){
        this.$router.go(-1)
    }
  },
  async created() {
   // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if(this.$route.query.id){
    const res = await findHosts({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.rehosts
       this.type = "update"
     }
    }else{
     this.type = "create"
   }
  
}
};
</script>

<style>
</style>