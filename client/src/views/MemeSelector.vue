<template>
  <div class="flex flex-col items-center w-full h-full">

    <div class="flex flex-col max-w-6xl px-5 my-10 w-100">
      <h1 class="text-3xl font-bold text-center text-blue-500">{{ emotion }}</h1>
      <p class="text-lg text-center text-gray-600">Selecione o meme que deseja usar.</p>
    </div>

    <div class="flex flex-wrap items-center justify-between max-w-6xl px-5 mt-10 w-100">
      <button
        class="w-64 h-auto mb-8 border-8 border-transparent border-solid shadow-xl hover:border-blue-500 transition-all duration-300 ease-out"
        v-for="image in memes"
        @click="() => select(image)"
        :key="image.content_url">
        <content-loader
          v-if="!loading[image.content_url]"
          height="280px"
          class="w-full" />
        <img
          :src="image.content_url"
          @load="() => toggleLoading(image.content_url)"
          :class="{ hidden: !loading[image.content_url] }"
          :alt="image.name"
          class="w-full"
        >
      </button>
    </div>
  </div>
</template>

<script>
import ContentLoader from '../components/ContentLoader'

export default {
  components: { ContentLoader },
  data () {
    return {
      emotion: this.$store.state.emotion,
      memes: this.$store.state.memes,
      loading: {}
    }
  },
  methods: {
    toggleLoading (key) {
      this.loading[key] = true
      this.$forceUpdate()
    },
    select (image) {
      this.$store.commit('setMeme', image.content_url)
      this.$router.push({ name: 'Result' })
    }
  }
}
</script>
