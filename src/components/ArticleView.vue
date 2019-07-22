<template>
  <div class="container">
    <div class="row">
      <div id="article-title" class="col-md-2">
        <router-link :to="`/`" tag="button" class="btn btn-primary btn-block pull-right">Back to Index</router-link>
      </div>
      <div class="col-md-2 offset-md-8">
        <router-link :to="`/edit/${articleTitle}`" tag="button" class="btn btn-primary btn-block pull-right">Edit</router-link>
      </div>
    </div>
    <div class="row">
      <div id="article-title" class="col-md-12">
        <h1>{{ articleTitle }}</h1>
      </div>
    </div>
    <hr />
    <div class="row">
      <div id="article-content" class="col-md-10 offset-md-1 text-left">
        <markdown-it-vue class="md-body" :content=articleContent />
      </div>
    </div>
  </div>
</template>

<script>
import MarkdownItVue from 'markdown-it-vue'

export default{
  components: {
    MarkdownItVue
  },
  data () {
    return {
      articleTitle: '',
      articleContent: ''
    }
  },
  methods: {

  },
  created () {
    this.$http.get('http://localhost:9090/articles/' + this.$route.params.name).then(function (data) {
      var articleJSON = data.body
      this.articleTitle = articleJSON['title']
      this.articleContent = articleJSON['content']
      // eslint-disable-next-line
    }, error => {
      this.$router.push('/')
    })
  }
}
</script>
