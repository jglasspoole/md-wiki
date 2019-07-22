import VueResource from 'vue-resource'
import Router from 'vue-router'
import IndexPage from '@/components/IndexPage.vue'
import { createLocalVue, shallowMount } from '@vue/test-utils'

const localVue = createLocalVue()
localVue.use(VueResource)
localVue.use(Router)

describe('IndexPage', () => {
  const router = new Router()
  const wrapper = shallowMount(IndexPage, {localVue, router})

  it('Confirms Vue instance', () => {
    expect(wrapper.isVueInstance()).toBe(true)
  })

  it('Renders the correct markup', () => {
    expect(wrapper.find('#index-title').exists()).toBe(true)
    expect(wrapper.find('#article-links').exists()).toBe(true)
  })
})