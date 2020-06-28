<template>
  <span
    :style="{ height, width: computedWidth }"
    class="opacity-75 content-loader"
  />
</template>

<script>
export default {
  props: {
    maxWidth: {
      default: 100,
      type: Number
    },
    minWidth: {
      default: 80,
      type: Number
    },
    height: {
      default: '1rem',
      type: String
    },
    width: {
      default: null,
      type: String
    }
  },
  computed: {
    computedWidth () {
      return this.width || `${Math.floor((Math.random() * (this.maxWidth - this.minWidth)) + this.minWidth)}%`
    }
  }
}
</script>

<style scoped>
@keyframes shimmer {
  100% {
    transform: translateX(100%);
  }
}
.content-loader {
  display: inline-block;
  position: relative;
  vertical-align: middle;
  overflow: hidden;
  background: #f6f7f8;
}
.content-loader::after {
  position: absolute;
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
  transform: translateX(-100%);
  background: #f6f7f8;
  background-image: linear-gradient(to right, #eeeeee 8%, #dddddd 18%, #eeeeee 33%);
  background-position: 0 0;
  background-size: 1000 100;
  animation-duration: 1.6s;
  animation: shimmer 1.6s infinite ease-in-out;
  content: '';
}
</style>
