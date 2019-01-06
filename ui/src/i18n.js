import 'moment/locale/en-gb'
import 'moment/locale/zh-cn'
import 'moment/locale/ja'
import moment from 'moment'

const allLocales = ['en', 'zh', 'ja']

let locale = navigator.languages && navigator.languages[0] || navigator.language || navigator.userLanguage || 'en'

locale = (locale.split('-')[0] || 'en').toLowerCase()

if (allLocales.indexOf(locale) == -1) {
    locale = 'en'
}

console.log('locale:', locale)

const momentLocalesMap = { en: 'en-gb', zh: 'zh-cn', ja: 'ja' }

moment.locale(momentLocalesMap[locale])

const messages = {
    en: {
        intro: {
            name: 'Canhead (罐头人 / 缶ヘッド)'
        },
        donation: {
            title: 'donation',
            button: 'donation'
        }
    },
    zh: {
        intro: {
            name: '罐头人 (Canhead / 缶ヘッド)'
        },
        donation: {
            title: '赞助',
            button: '赞助'
        }
    },
    ja: {
        intro: {
            name: '缶ヘッド (罐头人 / Canhead)'
        },
        donation: {
            title: '寄贈',
            button: '寄贈'
        }
    }
}

import Vue from 'vue'
import VueI18n from 'vue-i18n'

Vue.use(VueI18n)

export default new VueI18n({
    locale,
    fallbackLocale: 'en',
    messages
})
