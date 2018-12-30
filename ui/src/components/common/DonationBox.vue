<template>
    <b-card class="border-info">
        <b class="card-title"><v-icon name="money-bill" scale="1.4"></v-icon> Donation</b>
        <hr/>
        <div class="card-text">
            <p class="mb-0">
                <b>RMB</b> <small class="text-muted"> (wechat)</small>
            </p>
            <p class="text-center">
                <img class="rounded" src="/img/wx-donate-slice.jpg" style="width: 160px"/>
            </p>
            <p class="mb-0">
                <b>JPY / USD</b> <small class="text-muted"> (buymeacoffee.com)</small>
            </p>
            <p class="text-center mt-3 mb-0">
                <ch-bmcbutton></ch-bmcbutton>
            </p>
            <p class="mb-0">
                <b>Thanks</b>
            </p>
            <p v-if="sponsors.length > 0" class="mb-0">
                <span v-for="sponsor in sponsors" :key="sponsor.name">{{ sponsor.name }}</span>
            </p>
            <p v-if="sponsors.length == 0" class="mt-1 mb-1 text-center text-muted">
                no supporters yet
            </p>
        </div>
    </b-card>
</template>

<script>
    import BMCButton from './BMCButton.vue'

    export default {
        data () {
            return {
                sponsors: []
            }
        },
        components: {
            'ch-bmcbutton': BMCButton
        }
    }
</script>

<style>
span.sponsor {
}
</style>
