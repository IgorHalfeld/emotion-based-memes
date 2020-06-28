<template>
  <div class="flex flex-col items-center w-full h-full">
    <div class="flex flex-col items-center px-5">

      <loader v-if="!loaded" class="my-32" />
      <web-cam
        ref="camera"
        :class="{ hidden: !loaded }"
        class="mt-10 border-gray-500 border-solid shadow-xl border-5"
        @stopped="logErrors"
        @video-live="init"
        @error="logErrors"
        @cameras="loadCameraIds"
        :device-id="deviceId"
        :style="{
          width: '100%',
          maxWidth: '640px'
        }"
      />

      <p class="mt-3 italic text-gray-600">
        Tire uma foto para o Azure analizar e procurar um meme que combine com ela.
      </p>

      <button
        @click="take"
        class="w-full px-3 py-2 mt-10 text-2xl font-black text-white bg-blue-500 rounded-lg lg:w-auto focus:outline-none"
      >
        Take 🔥
      </button>
    </div>
  </div>
</template>

<script>
import { WebCam } from 'vue-web-cam'
import Loader from '../components/Loader'

export default {
  components: { WebCam, Loader },
  data: () => ({
    loaded: false,
    deviceId: null
  }),
  methods: {
    init () {
      this.loaded = true
    },
    loadCameraIds (device) {
      this.deviceId = device[0].deviceId
    },
    logErrors (error) {
      console.log('Error', error)
    },
    take () {
      this.$nextTick(async () => {
        try {
          const photo = this.$refs.camera.capture()
          this.$store.commit('setPicture', photo)
          this.$router.push({ name: 'Processing' })
        } catch (error) {
          console.log('Something happen', error)
        }
      })
    }
  }
}
</script>
