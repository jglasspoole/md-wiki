import VueResource from 'vue-resource'
import Router from 'vue-router'
import ArticleEdit from '@/components/ArticleEdit.vue'
import { createLocalVue, shallowMount } from '@vue/test-utils'

const localVue = createLocalVue()
localVue.use(VueResource)
localVue.use(Router)

describe('ArticleEdit', () => {
  const router = new Router()
  const wrapper = shallowMount(ArticleEdit, {localVue, router})
  
  it('Confirms Vue instance', () => {
    expect(wrapper.isVueInstance()).toBe(true)
  })

  it('Renders the correct markup', () => {
    expect(wrapper.find('#article-title').exists()).toBe(true)
    expect(wrapper.find('#article-content').exists()).toBe(true)
  })
})
