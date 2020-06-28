<template>
  <div class="flex flex-col loading">
    <loader />
    <p class="mt-8 text-gray-600">
      Procurando por um meme com o sentimento da foto..
    </p>
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
  created () {
    if (!this.$store.state.picture) {
      this.$router.push({ name: 'TakePicture' })
      return
    }

    this.handleUpload()
  },
  methods: {
    async handleUpload () {
      const file = dataURIToBlob(this.$store.state.picture)
      const form = new FormData()
      form.append('file', file, 'image.jpg')

      const res = await fetch('http://localhost:3000/face/analyze', { method: 'POST', body: form })
      const data = await res.json()

      this.$store.commit('setMemes', data.images)
      this.$store.commit('setEmotion', data.emotion)

      this.$router.push({ name: 'MemeSelector' })
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
