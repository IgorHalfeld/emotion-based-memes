<template>
  <div class="flex flex-col loading">
    <template v-if="isLoading">
      <loader />
      <p class="mt-8 text-gray-600">
        Criando o meme...
      </p>
    </template>
    <div
      v-else
      class="flex items-center justify-center w-full h-full max-w-6xl px-5 mt-10">
      <img :src="`http://localhost:3000/static/output.png?v=${Date.now()}`" alt="">
    </div>
  </div>
</template>

<script>
import Loader from '../components/Loader'

function dataURIToBlob (dataURI) {
  const splitDataURI = dataURI.split(',')
  const byteString = splitDataURI[0].indexOf('base64') >= 0
    ? atob(splitDataURI[1])
    : decodeURI(splitDataURI[1])

  const mimeString = splitDataURI[0].split(':')[1].split(';')[0]

  const ia = new Uint8Array(byteString.length)
  for (let i = 0; i < byteString.length; i++) {
    ia[i] = byteString.charCodeAt(i)
  }

  return new Blob([ia], { type: mimeString })
}

export default {
  components: { Loader },
  data: () => ({
    isLoading: false
  }),
  created () {
    this.handleUpload()
  },
  methods: {
    async handleUpload () {
      this.isLoading = true
      const file = dataURIToBlob(this.$store.state.picture)
      const form = new FormData()
      form.append('file', file, 'image.jpg')
      form.append('url', this.$store.state.meme)

      await fetch('http://localhost:3000/meme', { method: 'POST', body: form })

      this.isLoading = false
    }
  }
}
</script>

<style scoped>
@keyframes spin {
  to {
    transform: rotate(360deg);
  }
}

.loading {
  position: absolute;
  z-index: 9999;
  width: 100%;
  height: 80%;
  display: flex;
  background-color: #fff;
  justify-content: center;
  align-items: center;
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
