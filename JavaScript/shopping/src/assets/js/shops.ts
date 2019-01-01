import Vue from 'vue';

// 商品列表数据 - 虚拟 - 正式的情况下读取数据库，返回页面json格式
var products = {
    'p000001': { number: 'p000001', name: '魅族16', price: 2698, unit: '台', inventory: 1000, picture: 'meizu16.jpg' },
    'p000002': { number: 'p000002', name: '小米8', price: 2899, unit: '台', inventory: 200, picture: 'xiaomi8.jpg' },
    'p000003': { number: 'p000003', name: '华为P20', price: 3799, unit: '台', inventory: 3, picture: 'huaweip20.jpg' }
};

// 登录的用户 - 虚拟 - 正式的情况下获取登录的cookie或session值
var loginUser = 'wangshaocong';

/*------------------ 商店列表 --------------*/
// 组件-库存
var inventory = {};

var frontDesk = new Vue({
    el: '#frontDesk',
    data: {
        products: products
    },
    components: {
        'inventory': inventory,
    }
    // methods: {
        // 添加购物车方法
        // AddCart: function (id) {
        //     let addStatus = shopsStorage.addShoppingCart(id);
        //     if (!addStatus.status) {
        //         alert(addStatus.msg);
        //     }
        //     console.log(shopsStorage.getShoppingCart());
        // }
    // }
});
