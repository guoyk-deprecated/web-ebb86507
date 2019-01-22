<template>
<b-container>
    <ch-intro></ch-intro>
    <b-row>
        <b-col md="8" class="mt-2">
            <b-row>
                <b-col class="mb-3" v-for="post in posts" :key="post.ID" md="12">
                    <b-card class="shadow-sm" border-variant="light" header-bg-variant="light">
                        <p class="text-primary post-content" v-if="post.Content" v-html="post.Content"></p>
                        <p v-if="post.ImageURL" v-viewer>
                            <img class="post-image rounded" v-for="imageUrl in post.ImageURLs" :key="imageUrl" :src="imageUrl" />
                        </p>
                        <p class="mb-0">
                            <span class="text-muted">{{post.CreatedAt}}</span>
                        </p>
                        <div @click="onVoteClicked(post)" class="vote-area">
                            <div class="vote-heart" v-if="!post.Voted">
                                <v-icon class="text-danger" name="regular/heart" scale="1.2"></v-icon>
                            </div>
                            <div class="vote-heart text-danger" v-if="post.Voted">
                                <v-icon name="heart" scale="1.2"></v-icon>
                            </div>
                            <div class="vote-count-span">&nbsp;&nbsp;<span :class="post.Voted ? 'text-danger' : ''">{{post.VotesCount}}</span></div>
                        </div>
                        <a class="link-button" v-if="post.Link" :href="post.Link" target="_blank">
                            <v-icon name="link" scale="1.2"></v-icon>&nbsp;&nbsp;
                            <v-icon name="chevron-right" scale="1.2"></v-icon>
                        </a>
                    </b-card>
                </b-col>
                <b-col v-if="loading" md="12">
                    <b-card border-variant="light" header-bg-variant="light">
                        <p class="text-muted text-center mb-0">
                            <b><v-icon name="spinner" spin></v-icon> loading...</b>
                        </p>
                    </b-card>
                </b-col>
                <b-col v-if="hasMore && !loading" md="12">
                    <b-card @click="updatePosts" class="load-more-card" border-variant="light" header-bg-variant="light">
                        <p class="text-center mb-0">
                            <b><v-icon name="chevron-down"></v-icon> load more</b>
                        </p>
                    </b-card>
                </b-col>
            </b-row>
        </b-col>
        <b-col md="4" class="mt-2">
            <ch-donation></ch-donation>
        </b-col>
    </b-row>
    <b-row>
        <b-col class="text-center pt-4 pb-4">
            <small class="text-muted">2019 (c) canhead.xyz</small>
        </b-col>
    </b-row>
</b-container>
</template>

<script>
import IntroBox from './common/IntroBox.vue'
import DonationBox from './common/DonationBox.vue'
import moment from 'moment'

export default {
    name: 'index',
    data () {
        return {
            loading: false,
            hasMore: false,
            posts: []
        }
    },
    created () {
        this.$ga.page('/')
        this.updatePosts()
    },
    components: {
        'ch-intro': IntroBox,
        'ch-donation': DonationBox
    },
    methods: {
        updatePosts () {
            var params = {}
            if (this.posts.length > 0) {
                params.lastId = this.posts[this.posts.length - 1].ID
            }
            this.loading = true
            this.$http.get("/api/posts", { params }).then((res) => {
                let newPosts = res.body
                newPosts.forEach(post => {
                    post.ImageURLs = post.ImageURL.split('\n')
                    post.CreatedAt = moment(post.CreatedAt).local().calendar()
                })
                this.posts = this.posts.concat(newPosts)
                this.hasMore = newPosts.length == 50
                this.loading = false
            }, () => {
                this.loading = false
            })
        },
        onVoteClicked (post) {
            if (post.Voted) {
                return
            }
            post.Voted = true
            this.$http.post("/api/posts/" + post.ID + "/vote").then((res) => {
               post.VotesCount = res.body.VotesCount
            })
        }
    }
}
</script>

<style>
div.load-more-card {
    cursor: pointer;
}
img.post-image {
    height: 80px;
    margin-right: 0.8rem;
    margin-bottom: 0.6rem;
}
p.post-content {
    font-size: 1rem;
}
a.link-button {
    text-decoration: none;
    color: white;
    background-color: #18BC9C;
    position: absolute;
    right: -2px;
    bottom: 14px;
    width: 72px;
    height: 36px;
    cursor: pointer;
    text-align: center;
    padding-left: 8px;
    padding-top: 6px;
    border-top-left-radius: 0.25rem;
    border-bottom-left-radius: 0.25rem;
}
a.link-button:hover {
    color: white;
    background-color: #1acba8;
    -webkit-box-shadow: 0 0.125rem 0.25rem rgba(0, 0, 0, 0.075) !important;
    box-shadow: 0 0.125rem 0.25rem rgba(0, 0, 0, 0.075) !important;
}
div.vote-area {
    text-align: left;
    cursor: pointer;
    position: absolute;
    right: 82px;
    bottom: 16px;
    width: 60px;
    height: 32px;
}
div.vote-heart {
    position: absolute;
    top: 5px;
    left: 5px;
}
div.vote-count-span {
    position: absolute;
    left: 23px;
    top: 5px;
    width: 42px;
    font-size: 1rem;
}
</style>
