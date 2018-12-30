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
            <p class="text-center mt-3">
                <ch-bmcbutton></ch-bmcbutton>
            </p>
            <p>
                <b>Thanks</b>
            </p>
            <p v-if="sponsors.length > 0" class="mb-0">
                <span class="text-info sponsor" v-for="sponsor in sponsors" :key="sponsor.ID"><b><v-icon name="user" scale="0.8"></v-icon>&nbsp;{{ sponsor.Name }}</b></span>
            </p>
            <p v-if="sponsors.length == 0 && loaded" class="mt-1 mb-1 text-center text-muted">
                no supporters yet
            </p>
            <p v-if="sponsors.length == 0 && !loaded" class="mt-1 mb-1 text-center text-muted">
                loading...
            </p>
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
span.sponsor {
    margin-right: 0.8rem;
}
</style>
