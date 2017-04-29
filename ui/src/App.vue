<template>
  <main id="App">
    <h2 id="brand" class="flex-no-grow">
      SandBox Dashboard
    </h2>
    <section id="hero" class="flex-no-grow">
      <figure id="AppHero">
        20
        <figcaption>Hosts</figcaption>
      </figure>
      <figure id="RequestHero">
        1,223
        <figcaption>Request</figcaption>
      </figure>
      <figure id="SuccessHero">
        843
        <figcaption>Success</figcaption>
      </figure>
      <figure id="ErrorsHero">
        265
        <figcaption>Errors</figcaption>
      </figure>
    </section>
    <div id="target-list" class="panel flex-expand">
      <div class="panel-header">
        <nav>
          <div class="has-icon-left">
            <input class="form-input" placeholder="Search..." />
            <i class="form-icon icon icon-search"></i>
          </div>
          <button class="btn btn-link"><i class="icon icon-plus"></i> New Application</button>
        </nav>
      </div>
      <!-- Example of sorting DOM elements in JS -->
      <!-- https://github.com/uikit/uikit/blob/develop/src/js/components/sortable.js -->
      <!-- https://github.com/uikit/uikit/blob/develop/src/less/components/sortable.less -->
      <!-- No jQuery example: -->
      <!-- https://github.com/RubaXa/Sortable -->
      <!-- https://github.com/sagalbot/vue-sortable -->
      <div id="target-config-items" class="panel-body">
        <figure class="target-config-item--is-inactive">
          <h4> SandBox </h4>

          <i class="icon icon-shutdown"></i>
          <i class="icon icon-more-vert"></i>
        </figure>
        <figure class="target-config-item--is-active">
          <h4>MediaLocker</h4>

          <i class="icon icon-shutdown"></i>
          <i class="icon icon-more-vert"></i>
        </figure>
        <figure class="target-config-item--is-error">
          <h4>Homes</h4>

          <i class="icon icon-shutdown"></i>
          <i class="icon icon-more-vert"></i>
        </figure>

        <!-- contents -->
      </div>
      <!-- <div class="panel-footer"> -->
      <!--   <\!-- buttons or inputs -\-> -->
      <!-- </div> -->
    </div>
    <!-- <section id="content"> -->
    <!--   <\!-- <router-view></router-view> -\-> -->
    <!-- </section> -->
  </main>
</template>

<script>
export default {
  name: 'app'
}
</script>

<style lang="less">
@import "./less/spectre.less";
@import "./less/spectre-icons.less";
@import "./less/spectre-exp.less";

html,
body {
 height: 100vh;
}

#target-config-items {
  padding: 0;

  & figure {
    position: relative;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: flex-start;
    width: 100%;
    margin: 0;
    padding: 1.5rem 1.5rem 1.5rem 2.5rem;
    border-bottom: solid 1px @gray-color;

    & i {
      color: @gray-color;
      font-size: 2.8rem;
      transition: all .2s ease;
    }

    & i:hover {
      color: @body-font-color;
    }

    & i:last-child:hover {
      cursor: move;
    }

    & i:last-child {
      position: absolute;
      left: 1rem;

      color: @gray-color-light;
      font-size: 2.8rem;
      width: .5rem;
    }
  }

  & figure:hover,
  & figure:active {
    transition: all .2s ease;
    .control-shadow();

    & i:last-child {
      color: @body-font-color;
    }
  }

  & figure > h4 {
    align-items: flex-end;
    margin: 0;
  }

  & .target-config-item--is-inactive {
    border-left: .4rem solid @control-color-warning;
  }

  & .target-config-item--is-active {
    border-left: .4rem solid @control-color-success;
  }

  & .target-config-item--is-error {
    border-left: .4rem solid @control-color-error;
  }
}

#target-list nav {
  display: flex;
  justify-content: space-between;
  width: 100%;
  & div {
    max-width: 24rem;
  }

  & .has-icon-left i {
    color: @primary-color;
  }

  & .form-input {
    color: @gray-color-light;

    &:focus {
      color: @body-font-color;
    }
  }

  & .form-input {
    background: transparent;
    border: @border-width solid @border-color;
    &:hover,
    &:focus {
      border-color: @primary-color;
      background: transparent;
      .control-shadow();
    }
    &::placeholder {
      color: @gray-color-light;
      background: transparent;
      border-bottom: 1px solid @primary-color;
    }
  }
}


#target-list .panel-header {
  display: flex;
  justify-content: space-between;
}

#target-list .panel-body {
  border-top: 1px solid @border-color;
}

#App {
  display: flex;
  flex-flow: column;
  justify-content: flex-start;
  max-width: 76rem;
  min-width: 32rem;
  width: 100%;
  height: 100vh;
  margin: 0 auto;
  padding:  @layout-padding-lg;
  background: rgb(45, 45, 45);
}

.flex-no-grow {
  flex: 0 auto;
}

.flex-expand {
  flex: 2;
}

#brand {
  text-align: center;
  width: 100%;
}

#hero {
  display: flex;
  align-items: first baseline;
  // padding: 0 6.4rem 0 6.4rem;
  padding: 0 1rem 0 1rem;
  justify-content: space-around;
  // background: @hero-bg;
  background: rgba(45, 45, 45, .8);
  color: #FFF;
  padding: .8rem 0 1.2rem 0;
  margin-bottom: 1.5rem;
}

#hero figure {
  border-radius: 2px;
  padding: 1rem 2rem 1rem 2rem;
  margin: 0;
  min-width: 9.6rem;
  width: auto;
  line-height: 1.2;
  text-align: center;
  font-size: 4rem;
  font-weight: 400;
  color: rgba(255, 255, 255, .9);
  box-shadow: rgba(0, 0, 0, 0.0470588) 0px 1px 1px 0px;
}

#hero figcaption {
  font-variant-caps: all-small-caps;
  border-top: 1px dashed rgba(255, 255, 255, .5);
  font-size: 1.6rem;
  letter-spacing: 0.1rem;
}

#AppHero {
  background: rgb(44, 151, 222);
}

#RequestHero {
  background: rgb(231, 109, 59);
}

#SuccessHero {
  background: rgb(132, 181, 71);
}

#ErrorsHero {
  background: rgb(204, 62, 74);
}
</style>
