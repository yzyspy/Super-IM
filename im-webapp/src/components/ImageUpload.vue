<script setup lang="ts">
import { ref } from 'vue';
import  {uploadImageApi} from "@/api/file_api"


interface ImageUploadProps {
  image_type: string;
  extra?: string[]; // 可选属性
}

// 使用 defineProps 并指定泛型类型，同时通过 withDefaults 设置可选属性的默认值
const props = withDefaults(defineProps<ImageUploadProps>(), {
  extra: () => [''] // 数组或对象的默认值需要通过函数返回
});


// 使用defineEmits来定义事件
const emit = defineEmits(['onUploadSuccess']);

const uploadDone = (url : string) => {
  // 触发事件，并传递数据
  emit('onUploadSuccess', url);
};


const file = ref(null);
const uploading = ref(false);

const handleFileChange = (event:any) => {
  const selectedFile = event.target.files[0];
  if (selectedFile) {
    file.value = selectedFile;
  }
};

const upload = async () => {
  if (!file.value) return;

  uploading.value = true;
  try {
    const response = await uploadImageApi(file.value, props.image_type);
    console.log('上传成功:', response);
    uploadDone("http://localhost" + response.url )
    // 给父组件发送事件，通知图片上传成功，并且返回图片地址
  } catch (error) {
    console.error('上传失败:', error);
  } finally {
    uploading.value = false;
  }
};
</script>

<template>
  <div>{{image_type}}
    <input type="file" accept="image/*" @change="handleFileChange"/>
    <button @click="upload" :disabled="!file || uploading">
      {{ uploading ? '上传中...' : '上传图片' }}
    </button>
  </div>
</template>

<style scoped>

</style>