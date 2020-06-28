<template>
  <span class="loading-spinner"></span>
</template>

<script>
export default {
}
</script>

<style scoped>
@keyframes spin {
  to {
    transform: rotate(360deg);
  }
}

.loading-spinner {
  border: 4px solid transparent;
  border-radius: 50%;
  border-left-color: #4299e1;
  border-right-color: #4299e1;
  width: 100px;
  height: 100px;
  animation: spin 1s linear infinite;
}
</style>
