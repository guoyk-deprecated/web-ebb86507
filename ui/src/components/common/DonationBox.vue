<template>
    <b-card class="shadow-sm" border-variant="light" header-bg-variant="light">
        <template slot="header">
            <b><v-icon name="money-bill-wave" scale="1.2"></v-icon>&nbsp;&nbsp;{{ $t("donation.title") }}</b>
        </template>
        <div class="card-text">
            <p class="mb-1">
                <b><v-icon name="yen-sign"></v-icon> RMB</b> <small class="text-muted"> (wechat)</small>
            </p>
            <p class="text-center">
                <img class="rounded" src="../../assets/wx-donate.png" style="width: 160px"/>
            </p>
            <p class="mb-1">
                <b><v-icon name="dollar-sign"></v-icon> USD</b> <small class="text-muted"> (buymeacoffee.com)</small>
            </p>
            <p class="text-center mt-3">
                <ch-bmcbutton></ch-bmcbutton>
            </p>
            <p class="mb-2">
                <b><v-icon name="trophy"></v-icon> {{ $("donation.thanks") }} </b>
            </p>
            <b-row v-if="sponsors.length > 0">
                <b-col mb="12" v-for="sponsor in sponsors" :key="sponsor.ID" class="text-primary mb-0">
                    <v-icon name="award"></v-icon> {{ sponsor.Name }}
                </b-col>
            </b-row>
            <b-row v-if="sponsors.length == 0 && !loading">
                <b-col class="mt-1 mb-1 text-center text-muted">no supporters yet</b-col>
            </b-row>
            <b-row v-if="sponsors.length == 0 && loading">
                <b-col class="mt-1 mb-1 text-center text-muted">loading...</b-col>
            </b-row>
        </div>
    </b-card>
</template>

<script>
    import BMCButton from './BMCButton.vue'

    export default {
        data () {
            return {
                loading: false,
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
                this.loading = true
                this.$http.get("/api/sponsors").then(res=> {
                    this.sponsors = res.body
                    this.loading = false
                })
            }
        }
    }
</script>

<style>
.card-header {
    border-bottom: none;
}
</style>
