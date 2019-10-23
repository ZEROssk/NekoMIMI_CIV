var __extends = this.__extends || function (d, b) {
    for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p];
    function __() { this.constructor = d; }
    __.prototype = b.prototype;
    d.prototype = new __();
};
var util;
(function (util) {
    var Router = (function () {
        function Router(map) {
            this.map = map;
        }
        Router.prototype.getController = function () {
            if (this.map[location.pathname]) {
                return new this.map[location.pathname]();
            }
            return null;
        };
        return Router;
    })();
    util.Router = Router;

    var GetRouter = (function (_super) {
        __extends(GetRouter, _super);
        function GetRouter() {
            _super.apply(this, arguments);
        }
        GetRouter.prototype.getController = function () {
            if (this.map[location.search]) {
                return new this.map[location.search]();
            }
            return null;
        };
        return GetRouter;
    })(Router);
    util.GetRouter = GetRouter;
})(util || (util = {}));

var controllers;
(function (controllers) {
    var IndexController = (function () {
        function IndexController() {
        }
        IndexController.prototype.showAlert = function () {
            alert("IndexController");
        };
        return IndexController;
    })();
    controllers.IndexController = IndexController;

    var Sample1Controller = (function () {
        function Sample1Controller() {
        }
        Sample1Controller.prototype.showAlert = function () {
            alert("Sample1Controller");
        };
        return Sample1Controller;
    })();
    controllers.Sample1Controller = Sample1Controller;

    var Sample2Controller = (function () {
        function Sample2Controller() {
        }
        Sample2Controller.prototype.showAlert = function () {
            alert("Sample2Controller");
        };
        return Sample2Controller;
    })();
    controllers.Sample2Controller = Sample2Controller;
})(controllers || (controllers = {}));

var router = new util.GetRouter({
	"?page=index" : controllers.IndexController,
	"?page=sample1" : controllers.Sample1Controller,
	"?page=sample2" : controllers.Sample2Controller
});
var controller = router.getController();
controller.showAlert();

