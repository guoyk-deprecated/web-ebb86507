<template>
<b-container>
    <ch-intro></ch-intro>
    <b-row>
        <b-col>
            <b-card header="Admin Token">
                <b-form-group label="Admin Token:">
                    <b-form-input 
                      type="password"
                      v-model="adminToken"
                      required
                      placeholder="Enter admin token">
                    </b-form-input>
                </b-form-group>
            </b-card>
        </b-col>
    </b-row>
    <b-row class="mt-2">
        <b-col md="4">
            <b-card header="Add Sponsor">
                <b-form @submit.prevent="submitAddSponsorForm">
                    <b-form-group label="Name:">
                        <b-form-input 
                        type="text"
                        v-model="addSponsorForm.name"
                        required
                        placeholder="Enter Name">
                        </b-form-input>
                    </b-form-group>
                    <b-button type="submit" variant="primary">Submit</b-button>
                    <b-form-text>{{addSponsorForm.status}}</b-form-text>
                </b-form>
            </b-card>
        </b-col>
        <b-col md="4">
            <b-card header="Add Post">
                <b-form @submit.prevent="submitAddPostForm">
                    <b-form-group label="Content:">
                        <b-form-textarea
                        v-model="addPostForm.content"
                        placeholder="Enter Content">
                        </b-form-textarea>
                    </b-form-group>
                    <b-form-group label="Image URLs:">
                        <b-form-textarea
                        v-model="addPostForm.imageUrl"
                        placeholder="Enter Image URLs">
                        </b-form-textarea>
                    </b-form-group>
                    <b-button type="submit" variant="primary">Submit</b-button>
                    <b-form-text>{{addPostForm.status}}</b-form-text>
                </b-form>
            </b-card>
        </b-col>
        <b-col md="4">
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

export default {
    name: 'index',
    data () {
        return {
            adminToken: '',
            addSponsorForm: {
                name: '',
                status: '',
            },
            addPostForm: {
                content: '',
                imageUrl: '',
                status: ''
            }
        }
    },
    created () {
        this.adminToken = localStorage.getItem('admin-token')
    },
    components: {
        'ch-intro': IntroBox,
    },
    methods: {
        submitAddSponsorForm () {
            this.addSponsorForm.status = ''
            this.$http.post(
                '/api/sponsors/add', { 
                    name: this.addSponsorForm.name 
                }, {
                    headers: {
                        'X-Admin-Token': this.adminToken
                    }
                }).then((res) => {
                    localStorage.setItem('admin-token', this.adminToken)
                    this.addSponsorForm.status = "success: " + res.bodyText
                }, (res) => {
                    this.addSponsorForm.status = "error: " + res.bodyText
                })
        },
        submitAddPostForm () {
            this.addPostForm.status = ''
            this.$http.post(
                '/api/posts/add', { 
                    content: this.addPostForm.content,
                    imageUrl: this.addPostForm.imageUrl
                }, {
                    headers: {
                        'X-Admin-Token': this.adminToken
                    }
                }).then((res) => {
                    localStorage.setItem('admin-token', this.adminToken)
                    this.addPostForm.status = "success: " + res.bodyText
                }, (res) => {
                    this.addPostForm.status = "error: " + res.bodyText
                })
        }

    }
}
</script>

<style>
</style>
