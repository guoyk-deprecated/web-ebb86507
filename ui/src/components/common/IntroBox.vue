<template>
    <b-jumbotron class="main-jumbotron shadow-sm mt-4 mb-3">
        <b-row>
            <b-col md="4" class="pt-3 text-right">
                <img class="rounded-circle border border-secondary" src="../../assets/avatar-round-2.png" style="width: 160px; height: 160px;">
            </b-col>
            <b-col md="8" class="pt-4">
                <b-row>
                    <b-col>
                        <h2><b class="text-white">{{ $t("intro.name")  }}</b></h2>
                    </b-col>
                </b-row>
                <b-row class="mt-2 mb-2">
                    <b-col>
                        <h4><a :href="emailLink" class="text-secondary text-decoration-none"><v-icon name="envelope" scale="1.4"></v-icon>&nbsp; {{ email }}</a></h4>
                    </b-col>
                </b-row>
                <b-row class="mt-2 mb-2">
                    <b-col>
                        <h3>
                            <a v-for="link in links" :key="link.icon" class="text-secondary intro-link" :href="link.url" target="_blank"><b><v-icon scale="1.6" :name="link.icon"></v-icon></b></a>
                        </h3>
                    </b-col>
                </b-row>
            </b-col>
        </b-row>
    </b-jumbotron>
</template>

<script>
export default {
    data () {
        return {
            email: 'hi@canhead.xyz',
            emailLink: 'mailto:hi@canhead.xyz',
            links: [
                {icon: "lock", url: "/123EDB85.asc.txt"},
                {icon: "play-circle", url: "https://space.bilibili.com/19523038"},
                {icon: "brands/weibo", url: "https://weibo.com/canhead"},
                {icon: "brands/twitter", url: "https://twitter.com/virtcanhead"},
                {icon: "brands/instagram", url: "https://www.instagram.com/virtcanhead/"},
                {icon: "brands/github", url: "https://github.com/virtcanhead"},
                {icon: "brands/steam-square", url: "https://steamcommunity.com/id/virtcanhead"},
            ]
        }
    }
}
</script>

<style>
div.main-jumbotron {
    background-color: #333333;
    background-image: url("../../assets/jumbotron-bg.jpg");
}

a.intro-link {
    margin-right: 0.8rem;
}

@media (max-width: 576px) {
    a.intro-link {
        margin-right: 0.4rem;
    }
}
</style>
