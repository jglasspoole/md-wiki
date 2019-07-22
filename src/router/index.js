import Vue from 'vue'
import Router from 'vue-router'
import IndexPage from '../components/IndexPage'
import ArticleView from '../components/ArticleView'
import ArticleEdit from '../components/ArticleEdit'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'IndexPage',
      component: IndexPage
    },
    {
      path: '/:name',
      name: 'ArticleView',
      component: ArticleView
    },
    {
      path: '/edit/:name',
      name: 'ArticleEdit',
      component: ArticleEdit
    }
  ]
})
