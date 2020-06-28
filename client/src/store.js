import Vue from 'vue'

function createStore ({ state, mutations }) {
  return {
    state: Vue.observable(state),
    commit (key, ...args) {
      mutations[key](state, ...args)
    }
  }
}

const store = createStore({
  state: {
    picture: null,
    memes: [],
    meme: null,
    emotion: null
  },
  mutations: {
    setPicture (state, picture) {
      state.picture = picture
    },
    setMemes (state, memes) {
      state.memes = memes
    },
    setMeme (state, meme) {
      state.meme = meme
    },
    setEmotion (state, emotion) {
      state.emotion = `${emotion.slice(0, 1).toUpperCase()}${emotion.slice(1).toLowerCase()}`
    },
    cleanStore (state) {
      state.emotion = null
      state.picture = null
      state.memes = []
    }
  }
})

export default store
