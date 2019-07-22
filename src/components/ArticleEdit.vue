<template>
  <div class="container">
    <form id="app" @submit="validateForm" action="" method="put">
    <div class="row">
      <div id="article-title" class="col-md-12">
        <h1>{{ articleTitle }}</h1>
      </div>
    </div>
    <div v-if="formErrors.length" class="row">
      <div class="alert alert-danger col-md-12" role="alert">
          <span v-for="error in formErrors" v-bind:key="error">{{ error }}</span>
      </div>
    </div>
    <div class="row">
      <div id="article-content" class="col-md-12">
        <textarea v-model="articleContent" class="form-control" cols="30" rows="10" placeholder="Enter article content..." ></textarea>
      </div>
    </div>
    <div class="row mt-2">
      <div class="col-md-2 offset-md-6">
        <button type="button" id="preview-button" class="btn btn-primary btn-block" data-toggle="modal" data-target="#html-preview">Preview</button>
      </div>
      <div class="col-md-2">
        <input type="submit" value="Save" class="btn btn-primary btn-block" />
      </div>
      <div class="col-md-2">
        <router-link :to="`/${articleTitle}`" tag="button" class="btn btn-primary btn-block">Cancel</router-link>
      </div>
    </div>
    </form>

    <!-- Article HTML preview modal -->
    <div id="html-preview" class="modal fade" tabindex="-1" role="dialog" aria-hidden="true">
      <div class="modal-dialog modal-lg" role="document">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="col-8 modal-title text-center offset-2">{{ articleTitle }}</h5>
            <button type="button" class="close" data-dismiss="modal" aria-label="Close">
              <span aria-hidden="true">&times;</span>
            </button>
          </div>
          <div class="modal-body text-left">
            <markdown-it-vue class="md-body" :content=articleContent />
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-secondary" data-dismiss="modal">Close</button>
          </div>
        </div>
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
      articleContent: '',
      formErrors: []
    }
  },
  methods: {
    validateForm: function (e) {
      e.preventDefault()
      this.formErrors = []

      if (this.articleContent === '') {
        // Error to the form
        this.formErrors.push('Please enter content for the article.')
      } else {
        this.$http.put('http://localhost:9090/articles/' + this.articleTitle, this.articleContent).then(response => {
          // PUT request success
          this.$router.push('/' + this.articleTitle)
          // eslint-disable-next-line
        }, error => {
          // PUT request has failed
          this.formErrors.push('Something went wrong updating the article.')
        })
      }
    }
  },
  created () {
    this.$http.get('http://localhost:9090/articles/' + this.$route.params.name).then(function (data) {
      var articleJson = data.body
      this.articleTitle = articleJson['title']
      this.articleContent = articleJson['content']
      // eslint-disable-next-line
    }, error => {
      // Article was not found, we use the URL parameter as the title for a new article
      this.articleTitle = this.$route.params.name
    })
  }
}
</script>
