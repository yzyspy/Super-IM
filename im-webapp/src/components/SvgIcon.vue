<template>
  <svg :class="svgClass" aria-hidden="true">
    <use :xlink:href="iconClassName" />
  </svg>
</template>
<script setup lang="ts">
import { computed } from 'vue';
const props = defineProps({
  iconName: {
    type: String,
    required: true
  },
  className: {
    type: String,
    default: ''
  },
});
// 图标在 iconfont 中的名字
const iconClassName = computed(()=>{
  return `#${props.iconName}`;
})
// 给图标添加上类名
const svgClass = computed(() => {
  if (props.className) {
    return `svg-icon ${props.className}`;
  }
  return 'svg-icon';
});
</script>
<style scoped>
.svg-icon {
  width: 2em;
  height: 2em;
  position: relative;
  vertical-align: -2px;
  fill: #969896;
}
</style>