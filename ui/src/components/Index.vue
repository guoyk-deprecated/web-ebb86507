<template>
<b-container>
    <ch-intro></ch-intro>
    <b-row>
        <b-col md="8" class="mt-2">
            <b-card border-variant="light" header-bg-variant="light">
                <b class="card-title">
                    <v-icon name="comments" scale="1.4"></v-icon> Lorem Ipsum
                </b>
                <hr/>
                <p class="card-text">
                 Lorem ipsum dolor sit amet, consectetur adipiscing elit. Fusce mattis fermentum ullamcorper. Donec porttitor urna in orci volutpat, non euismod leo placerat. Morbi semper quam et massa lobortis euismod. Aenean pellentesque risus et gravida lobortis. Nam feugiat nisl ac neque volutpat, in dapibus leo tincidunt. In tincidunt ligula urna, lobortis dignissim lorem mattis sagittis. Donec ac sapien auctor metus semper cursus. Maecenas ut nibh ac erat tincidunt sodales sed vel mi.
                 Mauris ut lectus ut magna placerat pretium ut nec libero. Nunc vitae lobortis massa. Vestibulum non velit viverra, fermentum enim et, auctor lacus. Sed euismod nisi metus, id tincidunt mauris euismod sit amet.
                </p>
            </b-card>
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

export default {
    name: 'index',
    components: {
        'ch-intro': IntroBox,
        'ch-donation': DonationBox
    }
}
</script>

<style>
</style>
