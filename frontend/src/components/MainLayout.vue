<template>
  <div class="main-layout">
    <div class="sider"></div>
    <section class="main-screen">
      <div class="bread-nav">Nav</div>
      <div class="main-content">
        <slot></slot>
      </div>
    </section>
  </div>
</template>

<style scoped lang="scss">
.main-layout {
  display: flex;
  min-height: 100vh;
  min-width: 100vw;
  background: #252525;

  * {
    color: #efefef;
  }

  .sider {
    width: 60px;
    padding: 0 16px;
    background: #191919;
  }

  .main-screen {
    flex: 1;
    .bread-nav {
      padding: 8px;
      height: 40px;
    }

    .main-content {
      min-height: calc(100% - 40px);
      background: #303030;
    }
  }
}
</style>
