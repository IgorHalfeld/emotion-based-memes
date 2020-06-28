<template>
  <div class="w-full h-full">
    <custom-header />

    <router-view />
  </div>
</template>

<script>
import CustomHeader from './components/CustomHeader'

export default {
  components: { CustomHeader }
}
</script>
