<template>
    <b-card border-variant="info">
        <b class="card-title"><v-icon name="money-bill" scale="1.4"></v-icon> Donation</b>
        <hr/>
        <div class="card-text">
            <p class="mb-1">
                <b>RMB</b> <small class="text-muted"> (wechat)</small>
            </p>
            <p class="text-center">
                <img class="rounded" src="/img/wx-donate-slice.jpg" style="width: 160px"/>
            </p>
            <p class="mb-1">
                <b>JPY / USD</b> <small class="text-muted"> (buymeacoffee.com)</small>
            </p>
            <p class="text-center mt-3">
                <ch-bmcbutton></ch-bmcbutton>
            </p>
            <p class="mb-1">
                <b>Thanks</b>
            </p>
            <b-row>
                <b-col md="6" v-if="sponsors.length > 0" v-for="sponsor in sponsors" :key="sponsor.ID" class="text-info mb-0">
                    <b><v-icon name="award" scale="1"></v-icon>&nbsp;{{ sponsor.Name }}&nbsp;</b>
                </b-col>
                <b-col v-if="sponsors.length == 0 && loaded" class="mt-1 mb-1 text-center text-muted">no supporters yet</b-col>
                <b-col v-if="sponsors.length == 0 && !loaded" class="mt-1 mb-1 text-center text-muted">loading...</b-col>
            </b-row>
        </div>
    </b-card>
</template>

<script>
    import BMCButton from './BMCButton.vue'

    export default {
        data () {
            return {
                loaded: false,
                sponsors: []
            }
        },
        components: {
            'ch-bmcbutton': BMCButton
        },
        created () {
            this.updateSponsors()
        },
        methods: {
            updateSponsors () {
                this.$http.get("/api/sponsors").then(res=> {
                    this.sponsors = res.body
                    this.loaded = true
                })
            }
        }
    }
</script>

<style>
</style>
