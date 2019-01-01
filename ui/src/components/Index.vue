<template>
<b-container>
    <ch-intro></ch-intro>
    <b-row>
        <b-col md="8" class="mt-2">
            <b-row>
                <b-col class="mb-3" v-for="post in posts" :key="post.ID" md="12">
                    <b-card border-variant="light" header-bg-variant="light">
                        <p class="card-text" v-if="post.Content">{{post.Content}}</p>
                        <p class="card-text" v-if="post.ImageURL" v-viewer>
                            <img class="post-image" v-for="imageUrl in post.ImageURLs" :key="imageUrl" :src="imageUrl" />
                        </p>
                        <p class="card-text text-muted">
                            <small>{{post.CreatedAt}}</small>
                        </p>
                    </b-card>
                </b-col>
                <b-col v-if="loading" md="12">
                    <b-card border-variant="light" header-bg-variant="light">
                        <p class="card-text text-muted">
                            <b>loading...</b>
                        </p>
                    </b-card>
                </b-col>
                <b-col class="load-more-card" v-if="hasMore && !loading" md="12">
                    <b-card @click="updatePosts" border-variant="light" header-bg-variant="light">
                        <p class="card-text">
                            <b>load more</b>
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
            <small class="text-muted">2018 (c) canhead.xyz</small>
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
        }
    }
}
</script>

<style>
.load-more-card {
    cursor: pointer;
}
img.post-image {
    height: 160px;
    margin-right: 0.8rem;
}
</style>
