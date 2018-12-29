<template>
<b-container>
    <ch-intro></ch-intro>
</b-container>
</template>

<script>
import CommonIntro from './common/Intro.vue'

export default {
    name: 'index',
    components: {
        'ch-intro': CommonIntro,
    }
}
</script>

<style>
</style>
