<template>
    <b-card class="border-info">
        <b class="card-title"><v-icon name="money-bill" scale="1.4"></v-icon> Donation</b>
        <hr/>
        <div class="card-text">
            <p class="mb-0">
                RMB <small class="text-muted">(wechat)</small>
            </p>
            <p class="text-center">
                <img class="rounded" src="/img/wx-donate-slice.jpg" style="width: 160px"/>
            </p>
            <p class="mb-0">
                JPY / USD <small class="text-muted">(buymeacoffee.com)</small>
            </p>
            <p class="text-center mt-3 mb-0">
                <ch-bmcbutton></ch-bmcbutton>
            </p>
        </div>
    </b-card>
</template>

<script>
    import BMCButton from './BMCButton.vue'

    export default {
        components: {
            'ch-bmcbutton': BMCButton
        }
    }
</script>

<style>
</style>
