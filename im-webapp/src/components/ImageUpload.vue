<script setup lang="ts">
import { ref } from 'vue';
import  {uploadImageApi} from "@/api/file_api"

const file = ref(null);
const uploading = ref(false);
const previewUrl = ref('');

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
    const response = await uploadImageApi(file.value, 'avatar');
    console.log('上传成功:', response);
    // 成功后的处理
  } catch (error) {
    console.error('上传失败:', error);
  } finally {
    uploading.value = false;
  }
};
</script>

<template>
  <div>
    <input
        type="file"
        alt="上传头像"
        accept="image/*"
        @change="handleFileChange"
    />
    <button @click="upload" :disabled="!file || uploading">
      {{ uploading ? '上传中...' : '上传图片' }}
    </button>
    <!-- 预览图片 -->
    <div v-if="previewUrl" class="preview">
      <img :src="previewUrl" alt="预览" style="max-width: 200px; max-height: 200px;" />
    </div>
  </div>
</template>

<style scoped>

</style>