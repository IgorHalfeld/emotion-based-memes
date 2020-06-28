<template>
  <div class="flex justify-center w-full h-32 bg-gray-200">
    <div class="fixed top-0 left-0 mt-8 ml-8" v-if="canShowCreateNew">
      <button
        @click="createNew"
        class="p-3 text-lg text-white bg-blue-500 rounded-full hover:bg-white hover:text-blue-500 transition-all ease-out duration-500">
        ← criar outro
      </button>
    </div>
    <div class="flex flex-col items-center justify-center max-w-xl">
      <h3 class="text-4xl font-black text-blue-500">
        Emotion-based Memes
      </h3>
      <span class="font-bold text-gray-600">
        Powered by Azure
      </span>
    </div>
  </div>
</template>

<script>
export default {
  computed: {
    canShowCreateNew () {
      return this.$route.name !== 'TakePicture'
    }
  },
  methods: {
    createNew () {
      this.$store.commit('cleanStore')
      this.$router.push({ name: 'TakePicture' })
    }
  }
}
</script>
