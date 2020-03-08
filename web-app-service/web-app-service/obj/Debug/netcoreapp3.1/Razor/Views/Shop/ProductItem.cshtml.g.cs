#pragma checksum "C:\source\repos\sneaker-resale-platform\web-app-service\web-app-service\Views\Shop\ProductItem.cshtml" "{ff1816ec-aa5e-4d10-87f7-6f4963833460}" "448c260f191432df0b6b0c3275858c9e82cd4115"
// <auto-generated/>
#pragma warning disable 1591
[assembly: global::Microsoft.AspNetCore.Razor.Hosting.RazorCompiledItemAttribute(typeof(AspNetCore.Views_Shop_ProductItem), @"mvc.1.0.view", @"/Views/Shop/ProductItem.cshtml")]
namespace AspNetCore
{
    #line hidden
    using System;
    using System.Collections.Generic;
    using System.Linq;
    using System.Threading.Tasks;
    using Microsoft.AspNetCore.Mvc;
    using Microsoft.AspNetCore.Mvc.Rendering;
    using Microsoft.AspNetCore.Mvc.ViewFeatures;
#nullable restore
#line 1 "C:\source\repos\sneaker-resale-platform\web-app-service\web-app-service\Views\_ViewImports.cshtml"
using web_app_service;

#line default
#line hidden
#nullable disable
#nullable restore
#line 2 "C:\source\repos\sneaker-resale-platform\web-app-service\web-app-service\Views\_ViewImports.cshtml"
using web_app_service.Models;

#line default
#line hidden
#nullable disable
    [global::Microsoft.AspNetCore.Razor.Hosting.RazorSourceChecksumAttribute(@"SHA1", @"448c260f191432df0b6b0c3275858c9e82cd4115", @"/Views/Shop/ProductItem.cshtml")]
    [global::Microsoft.AspNetCore.Razor.Hosting.RazorSourceChecksumAttribute(@"SHA1", @"eb13faaba876e14f119efae96039076ffb6f688a", @"/Views/_ViewImports.cshtml")]
    public class Views_Shop_ProductItem : global::Microsoft.AspNetCore.Mvc.Razor.RazorPage<web_app_service.Models.SneakerProduct>
    {
        private static readonly global::Microsoft.AspNetCore.Razor.TagHelpers.TagHelperAttribute __tagHelperAttribute_0 = new global::Microsoft.AspNetCore.Razor.TagHelpers.TagHelperAttribute("href", new global::Microsoft.AspNetCore.Html.HtmlString("~/css/product_item.css"), global::Microsoft.AspNetCore.Razor.TagHelpers.HtmlAttributeValueStyle.DoubleQuotes);
        private static readonly global::Microsoft.AspNetCore.Razor.TagHelpers.TagHelperAttribute __tagHelperAttribute_1 = new global::Microsoft.AspNetCore.Razor.TagHelpers.TagHelperAttribute("rel", new global::Microsoft.AspNetCore.Html.HtmlString("stylesheet"), global::Microsoft.AspNetCore.Razor.TagHelpers.HtmlAttributeValueStyle.DoubleQuotes);
        private static readonly global::Microsoft.AspNetCore.Razor.TagHelpers.TagHelperAttribute __tagHelperAttribute_2 = new global::Microsoft.AspNetCore.Razor.TagHelpers.TagHelperAttribute("href", new global::Microsoft.AspNetCore.Html.HtmlString("~/css/product_item_responsive.css"), global::Microsoft.AspNetCore.Razor.TagHelpers.HtmlAttributeValueStyle.DoubleQuotes);
        private static readonly global::Microsoft.AspNetCore.Razor.TagHelpers.TagHelperAttribute __tagHelperAttribute_3 = new global::Microsoft.AspNetCore.Razor.TagHelpers.TagHelperAttribute("asp-area", "", global::Microsoft.AspNetCore.Razor.TagHelpers.HtmlAttributeValueStyle.DoubleQuotes);
        private static readonly global::Microsoft.AspNetCore.Razor.TagHelpers.TagHelperAttribute __tagHelperAttribute_4 = new global::Microsoft.AspNetCore.Razor.TagHelpers.TagHelperAttribute("asp-controller", "Shop", global::Microsoft.AspNetCore.Razor.TagHelpers.HtmlAttributeValueStyle.DoubleQuotes);
        private static readonly global::Microsoft.AspNetCore.Razor.TagHelpers.TagHelperAttribute __tagHelperAttribute_5 = new global::Microsoft.AspNetCore.Razor.TagHelpers.TagHelperAttribute("asp-action", "ProductItem", global::Microsoft.AspNetCore.Razor.TagHelpers.HtmlAttributeValueStyle.DoubleQuotes);
        private static readonly global::Microsoft.AspNetCore.Razor.TagHelpers.TagHelperAttribute __tagHelperAttribute_6 = new global::Microsoft.AspNetCore.Razor.TagHelpers.TagHelperAttribute("src", new global::Microsoft.AspNetCore.Html.HtmlString("~/js/productItem.js"), global::Microsoft.AspNetCore.Razor.TagHelpers.HtmlAttributeValueStyle.DoubleQuotes);
        #line hidden
        #pragma warning disable 0649
        private global::Microsoft.AspNetCore.Razor.Runtime.TagHelpers.TagHelperExecutionContext __tagHelperExecutionContext;
        #pragma warning restore 0649
        private global::Microsoft.AspNetCore.Razor.Runtime.TagHelpers.TagHelperRunner __tagHelperRunner = new global::Microsoft.AspNetCore.Razor.Runtime.TagHelpers.TagHelperRunner();
        #pragma warning disable 0169
        private string __tagHelperStringValueBuffer;
        #pragma warning restore 0169
        private global::Microsoft.AspNetCore.Razor.Runtime.TagHelpers.TagHelperScopeManager __backed__tagHelperScopeManager = null;
        private global::Microsoft.AspNetCore.Razor.Runtime.TagHelpers.TagHelperScopeManager __tagHelperScopeManager
        {
            get
            {
                if (__backed__tagHelperScopeManager == null)
                {
                    __backed__tagHelperScopeManager = new global::Microsoft.AspNetCore.Razor.Runtime.TagHelpers.TagHelperScopeManager(StartTagHelperWritingScope, EndTagHelperWritingScope);
                }
                return __backed__tagHelperScopeManager;
            }
        }
        private global::Microsoft.AspNetCore.Mvc.Razor.TagHelpers.UrlResolutionTagHelper __Microsoft_AspNetCore_Mvc_Razor_TagHelpers_UrlResolutionTagHelper;
        private global::Microsoft.AspNetCore.Mvc.TagHelpers.AnchorTagHelper __Microsoft_AspNetCore_Mvc_TagHelpers_AnchorTagHelper;
        #pragma warning disable 1998
        public async override global::System.Threading.Tasks.Task ExecuteAsync()
        {
#nullable restore
#line 2 "C:\source\repos\sneaker-resale-platform\web-app-service\web-app-service\Views\Shop\ProductItem.cshtml"
  
    Layout = "~/Views/Shared/_Layout.cshtml";
    var images = Model.Images.Select(image => "/images/" + image).ToList();
    var relatedProducts = ViewBag.RelatedProducts as List<SneakerProduct>;

#line default
#line hidden
#nullable disable
            WriteLiteral("\r\n");
            DefineSection("Styles", async() => {
                WriteLiteral("\r\n    ");
                __tagHelperExecutionContext = __tagHelperScopeManager.Begin("link", global::Microsoft.AspNetCore.Razor.TagHelpers.TagMode.SelfClosing, "448c260f191432df0b6b0c3275858c9e82cd41156166", async() => {
                }
                );
                __Microsoft_AspNetCore_Mvc_Razor_TagHelpers_UrlResolutionTagHelper = CreateTagHelper<global::Microsoft.AspNetCore.Mvc.Razor.TagHelpers.UrlResolutionTagHelper>();
                __tagHelperExecutionContext.Add(__Microsoft_AspNetCore_Mvc_Razor_TagHelpers_UrlResolutionTagHelper);
                __tagHelperExecutionContext.AddHtmlAttribute(__tagHelperAttribute_0);
                __tagHelperExecutionContext.AddHtmlAttribute(__tagHelperAttribute_1);
                await __tagHelperRunner.RunAsync(__tagHelperExecutionContext);
                if (!__tagHelperExecutionContext.Output.IsContentModified)
                {
                    await __tagHelperExecutionContext.SetOutputContentAsync();
                }
                Write(__tagHelperExecutionContext.Output);
                __tagHelperExecutionContext = __tagHelperScopeManager.End();
                WriteLiteral("\r\n    ");
                __tagHelperExecutionContext = __tagHelperScopeManager.Begin("link", global::Microsoft.AspNetCore.Razor.TagHelpers.TagMode.SelfClosing, "448c260f191432df0b6b0c3275858c9e82cd41157344", async() => {
                }
                );
                __Microsoft_AspNetCore_Mvc_Razor_TagHelpers_UrlResolutionTagHelper = CreateTagHelper<global::Microsoft.AspNetCore.Mvc.Razor.TagHelpers.UrlResolutionTagHelper>();
                __tagHelperExecutionContext.Add(__Microsoft_AspNetCore_Mvc_Razor_TagHelpers_UrlResolutionTagHelper);
                __tagHelperExecutionContext.AddHtmlAttribute(__tagHelperAttribute_2);
                __tagHelperExecutionContext.AddHtmlAttribute(__tagHelperAttribute_1);
                await __tagHelperRunner.RunAsync(__tagHelperExecutionContext);
                if (!__tagHelperExecutionContext.Output.IsContentModified)
                {
                    await __tagHelperExecutionContext.SetOutputContentAsync();
                }
                Write(__tagHelperExecutionContext.Output);
                __tagHelperExecutionContext = __tagHelperScopeManager.End();
                WriteLiteral("\r\n");
            }
            );
            WriteLiteral(@"
<!-- Product Details -->
<div class=""product_details"">
    <div class=""container"">
        <div class=""row details_row"">
            <!-- Product Image -->
            <div class=""col-lg-6"">
                <div class=""details_image"">
                    <div class=""details_image_large"">
                        <img");
            BeginWriteAttribute("src", " src=", 738, "", 772, 1);
#nullable restore
#line 22 "C:\source\repos\sneaker-resale-platform\web-app-service\web-app-service\Views\Shop\ProductItem.cshtml"
WriteAttributeValue("", 743, images.ElementAtOrDefault(0), 743, 29, false);

#line default
#line hidden
#nullable disable
            EndWriteAttribute();
            BeginWriteAttribute("alt", " alt=\"", 772, "\"", 778, 0);
            EndWriteAttribute();
            WriteLiteral(">\r\n");
#nullable restore
#line 23 "C:\source\repos\sneaker-resale-platform\web-app-service\web-app-service\Views\Shop\ProductItem.cshtml"
                         if (DateTime.Today.Subtract(Model.AddedAt).Days < 5)
                        {

#line default
#line hidden
#nullable disable
            WriteLiteral("                            <div class=\"product_extra product_new\"><a>New</a></div>\r\n");
#nullable restore
#line 26 "C:\source\repos\sneaker-resale-platform\web-app-service\web-app-service\Views\Shop\ProductItem.cshtml"
                        }

#line default
#line hidden
#nullable disable
            WriteLiteral("                    </div>\r\n                    <div class=\"details_image_thumbnails d-flex flex-row align-items-start justify-content-between\">\r\n");
#nullable restore
#line 29 "C:\source\repos\sneaker-resale-platform\web-app-service\web-app-service\Views\Shop\ProductItem.cshtml"
                         foreach (var image in images)
                        {

#line default
#line hidden
#nullable disable
            WriteLiteral("                            <div class=\"details_image_thumbnail\" data-image=");
#nullable restore
#line 31 "C:\source\repos\sneaker-resale-platform\web-app-service\web-app-service\Views\Shop\ProductItem.cshtml"
                                                                       Write(image);

#line default
#line hidden
#nullable disable
            WriteLiteral("><img");
            BeginWriteAttribute("src", " src=", 1316, "", 1327, 1);
#nullable restore
#line 31 "C:\source\repos\sneaker-resale-platform\web-app-service\web-app-service\Views\Shop\ProductItem.cshtml"
WriteAttributeValue("", 1321, image, 1321, 6, false);

#line default
#line hidden
#nullable disable
            EndWriteAttribute();
            BeginWriteAttribute("alt", " alt=\"", 1327, "\"", 1333, 0);
            EndWriteAttribute();
            WriteLiteral("></div>\r\n");
#nullable restore
#line 32 "C:\source\repos\sneaker-resale-platform\web-app-service\web-app-service\Views\Shop\ProductItem.cshtml"
                        }

#line default
#line hidden
#nullable disable
            WriteLiteral("                    </div>\r\n                </div>\r\n            </div>\r\n\r\n            <!-- Product Content -->\r\n            <div class=\"col-lg-6\">\r\n                <div class=\"details_content\">\r\n                    <div class=\"details_name\">");
#nullable restore
#line 40 "C:\source\repos\sneaker-resale-platform\web-app-service\web-app-service\Views\Shop\ProductItem.cshtml"
                                         Write(Html.DisplayFor(model => Model.BrandName));

#line default
#line hidden
#nullable disable
            WriteLiteral("</div>\r\n                    <div class=\"details_name\">");
#nullable restore
#line 41 "C:\source\repos\sneaker-resale-platform\web-app-service\web-app-service\Views\Shop\ProductItem.cshtml"
                                         Write(Html.DisplayFor(model => Model.ModelName));

#line default
#line hidden
#nullable disable
            WriteLiteral("</div>\r\n                    <div class=\"details_price\">");
#nullable restore
#line 42 "C:\source\repos\sneaker-resale-platform\web-app-service\web-app-service\Views\Shop\ProductItem.cshtml"
                                          Write(Html.DisplayFor(model => Model.Price));

#line default
#line hidden
#nullable disable
            WriteLiteral(@"</div>

                    <!-- In Stock -->
                    <div class=""in_stock_container"">
                        <div class=""availability"">Availability:</div>
                        <span>In Stock</span>
                    </div>
                    <div class=""details_text"">
                        <p>");
#nullable restore
#line 50 "C:\source\repos\sneaker-resale-platform\web-app-service\web-app-service\Views\Shop\ProductItem.cshtml"
                      Write(Model.Description);

#line default
#line hidden
#nullable disable
            WriteLiteral(@"</p>
                    </div>

                    <!-- Product Quantity -->
                    <div class=""product_quantity_container"">
                        <div class=""product_quantity clearfix"">
                            <span>Qty</span>
                            <input id=""quantity_input"" type=""text"" pattern=""[0-9]*"" value=""1"">
                            <div class=""quantity_buttons"">
                                <div id=""quantity_inc_button"" class=""quantity_inc quantity_control""><i class=""fa fa-chevron-up"" aria-hidden=""true""></i></div>
                                <div id=""quantity_dec_button"" class=""quantity_dec quantity_control""><i class=""fa fa-chevron-down"" aria-hidden=""true""></i></div>
                            </div>
                        </div>
                        <div class=""button cart_button""><a href=""#"">Add to cart</a></div>
                    </div>

                    <!-- Share -->
                    <div class=""details_share"">
                  ");
            WriteLiteral(@"      <span>Share:</span>
                        <ul>
                            <li><a href=""#""><i class=""fa fa-pinterest"" aria-hidden=""true""></i></a></li>
                            <li><a href=""#""><i class=""fa fa-instagram"" aria-hidden=""true""></i></a></li>
                            <li><a href=""#""><i class=""fa fa-facebook"" aria-hidden=""true""></i></a></li>
                            <li><a href=""#""><i class=""fa fa-twitter"" aria-hidden=""true""></i></a></li>
                        </ul>
                    </div>
                </div>
            </div>
        </div>

        <div class=""row description_row"">
            <div class=""col"">
                <div class=""description_title_container"">
                    <div class=""description_title"">Description</div>
                    <div class=""reviews_title""><a href=""#"">Reviews <span>(1)</span></a></div>
                </div>
                <div class=""description_text"">
                    <p>");
#nullable restore
#line 87 "C:\source\repos\sneaker-resale-platform\web-app-service\web-app-service\Views\Shop\ProductItem.cshtml"
                  Write(Model.Description);

#line default
#line hidden
#nullable disable
            WriteLiteral("</p>\r\n                </div>\r\n            </div>\r\n        </div>\r\n    </div>\r\n</div>\r\n\r\n<!-- Related Products -->\r\n");
#nullable restore
#line 95 "C:\source\repos\sneaker-resale-platform\web-app-service\web-app-service\Views\Shop\ProductItem.cshtml"
 if (relatedProducts != null && relatedProducts.Any())
{

#line default
#line hidden
#nullable disable
            WriteLiteral(@"    <div class=""products"">
        <div class=""container"">
            <div class=""row"">
                <div class=""col text-center"">
                    <div class=""products_title"">Related Products</div>
                </div>
            </div>
            <div class=""row"">
                <div class=""col"">
                    <div class=""product_grid"">
");
#nullable restore
#line 107 "C:\source\repos\sneaker-resale-platform\web-app-service\web-app-service\Views\Shop\ProductItem.cshtml"
                         for (var y = 0; y < @relatedProducts.Count; y++)
                        {

#line default
#line hidden
#nullable disable
            WriteLiteral("                            <div class=\"row\">\r\n");
#nullable restore
#line 110 "C:\source\repos\sneaker-resale-platform\web-app-service\web-app-service\Views\Shop\ProductItem.cshtml"
                                 for (var x = 0; x <= 4; x++)
                                {
                                    if (y >= @relatedProducts.Count)
                                    {
                                        break;
                                    }
                                    var product = @relatedProducts[y++];
                                    var image = "/images/" + product.Images.FirstOrDefault() ?? string.Empty;
                                    var price = $"{product.Price}$";
                                    

#line default
#line hidden
#nullable disable
            WriteLiteral("<!-- Product -->\r\n                                    <div class=\"col-sm product\">\r\n                                        <div class=\"product_image\"><img");
            BeginWriteAttribute("src", " src=", 5646, "", 5657, 1);
#nullable restore
#line 121 "C:\source\repos\sneaker-resale-platform\web-app-service\web-app-service\Views\Shop\ProductItem.cshtml"
WriteAttributeValue("", 5651, image, 5651, 6, false);

#line default
#line hidden
#nullable disable
            EndWriteAttribute();
            BeginWriteAttribute("alt", " alt=\"", 5657, "\"", 5663, 0);
            EndWriteAttribute();
            WriteLiteral("></div>\r\n");
#nullable restore
#line 122 "C:\source\repos\sneaker-resale-platform\web-app-service\web-app-service\Views\Shop\ProductItem.cshtml"
                                         if (DateTime.Today.Subtract(product.AddedAt).Days < 5)
                                        {

#line default
#line hidden
#nullable disable
            WriteLiteral("                                            <div class=\"product_extra product_new\"><a>New</a></div>\r\n");
#nullable restore
#line 125 "C:\source\repos\sneaker-resale-platform\web-app-service\web-app-service\Views\Shop\ProductItem.cshtml"
                                        }

#line default
#line hidden
#nullable disable
            WriteLiteral("                                        <div class=\"product_content\">\r\n                                            <div class=\"product_title\">");
            __tagHelperExecutionContext = __tagHelperScopeManager.Begin("a", global::Microsoft.AspNetCore.Razor.TagHelpers.TagMode.StartTagAndEndTag, "448c260f191432df0b6b0c3275858c9e82cd411519376", async() => {
#nullable restore
#line 127 "C:\source\repos\sneaker-resale-platform\web-app-service\web-app-service\Views\Shop\ProductItem.cshtml"
                                                                                                                                                                Write(Html.DisplayFor(modelItem => product.ModelName));

#line default
#line hidden
#nullable disable
            }
            );
            __Microsoft_AspNetCore_Mvc_TagHelpers_AnchorTagHelper = CreateTagHelper<global::Microsoft.AspNetCore.Mvc.TagHelpers.AnchorTagHelper>();
            __tagHelperExecutionContext.Add(__Microsoft_AspNetCore_Mvc_TagHelpers_AnchorTagHelper);
            __Microsoft_AspNetCore_Mvc_TagHelpers_AnchorTagHelper.Area = (string)__tagHelperAttribute_3.Value;
            __tagHelperExecutionContext.AddTagHelperAttribute(__tagHelperAttribute_3);
            __Microsoft_AspNetCore_Mvc_TagHelpers_AnchorTagHelper.Controller = (string)__tagHelperAttribute_4.Value;
            __tagHelperExecutionContext.AddTagHelperAttribute(__tagHelperAttribute_4);
            __Microsoft_AspNetCore_Mvc_TagHelpers_AnchorTagHelper.Action = (string)__tagHelperAttribute_5.Value;
            __tagHelperExecutionContext.AddTagHelperAttribute(__tagHelperAttribute_5);
            if (__Microsoft_AspNetCore_Mvc_TagHelpers_AnchorTagHelper.RouteValues == null)
            {
                throw new InvalidOperationException(InvalidTagHelperIndexerAssignment("asp-route-productId", "Microsoft.AspNetCore.Mvc.TagHelpers.AnchorTagHelper", "RouteValues"));
            }
            BeginWriteTagHelperAttribute();
#nullable restore
#line 127 "C:\source\repos\sneaker-resale-platform\web-app-service\web-app-service\Views\Shop\ProductItem.cshtml"
                                                                                                                                             WriteLiteral(product.Id);

#line default
#line hidden
#nullable disable
            __tagHelperStringValueBuffer = EndWriteTagHelperAttribute();
            __Microsoft_AspNetCore_Mvc_TagHelpers_AnchorTagHelper.RouteValues["productId"] = __tagHelperStringValueBuffer;
            __tagHelperExecutionContext.AddTagHelperAttribute("asp-route-productId", __Microsoft_AspNetCore_Mvc_TagHelpers_AnchorTagHelper.RouteValues["productId"], global::Microsoft.AspNetCore.Razor.TagHelpers.HtmlAttributeValueStyle.DoubleQuotes);
            await __tagHelperRunner.RunAsync(__tagHelperExecutionContext);
            if (!__tagHelperExecutionContext.Output.IsContentModified)
            {
                await __tagHelperExecutionContext.SetOutputContentAsync();
            }
            Write(__tagHelperExecutionContext.Output);
            __tagHelperExecutionContext = __tagHelperScopeManager.End();
            WriteLiteral("</div>\r\n                                            <div class=\"product_price\">");
#nullable restore
#line 128 "C:\source\repos\sneaker-resale-platform\web-app-service\web-app-service\Views\Shop\ProductItem.cshtml"
                                                                  Write(price);

#line default
#line hidden
#nullable disable
            WriteLiteral("</div>\r\n                                        </div>\r\n                                    </div>\r\n");
#nullable restore
#line 131 "C:\source\repos\sneaker-resale-platform\web-app-service\web-app-service\Views\Shop\ProductItem.cshtml"
                                }

#line default
#line hidden
#nullable disable
            WriteLiteral("                            </div>\r\n");
#nullable restore
#line 133 "C:\source\repos\sneaker-resale-platform\web-app-service\web-app-service\Views\Shop\ProductItem.cshtml"
                        }

#line default
#line hidden
#nullable disable
            WriteLiteral("                    </div>\r\n                </div>\r\n            </div>\r\n        </div>\r\n    </div>\r\n");
#nullable restore
#line 139 "C:\source\repos\sneaker-resale-platform\web-app-service\web-app-service\Views\Shop\ProductItem.cshtml"
}

#line default
#line hidden
#nullable disable
            WriteLiteral("\r\n\r\n");
            DefineSection("Scripts", async() => {
                WriteLiteral("\r\n    ");
                __tagHelperExecutionContext = __tagHelperScopeManager.Begin("script", global::Microsoft.AspNetCore.Razor.TagHelpers.TagMode.StartTagAndEndTag, "448c260f191432df0b6b0c3275858c9e82cd411523893", async() => {
                }
                );
                __Microsoft_AspNetCore_Mvc_Razor_TagHelpers_UrlResolutionTagHelper = CreateTagHelper<global::Microsoft.AspNetCore.Mvc.Razor.TagHelpers.UrlResolutionTagHelper>();
                __tagHelperExecutionContext.Add(__Microsoft_AspNetCore_Mvc_Razor_TagHelpers_UrlResolutionTagHelper);
                __tagHelperExecutionContext.AddHtmlAttribute(__tagHelperAttribute_6);
                await __tagHelperRunner.RunAsync(__tagHelperExecutionContext);
                if (!__tagHelperExecutionContext.Output.IsContentModified)
                {
                    await __tagHelperExecutionContext.SetOutputContentAsync();
                }
                Write(__tagHelperExecutionContext.Output);
                __tagHelperExecutionContext = __tagHelperScopeManager.End();
                WriteLiteral("\r\n");
            }
            );
        }
        #pragma warning restore 1998
        [global::Microsoft.AspNetCore.Mvc.Razor.Internal.RazorInjectAttribute]
        public global::Microsoft.AspNetCore.Mvc.ViewFeatures.IModelExpressionProvider ModelExpressionProvider { get; private set; }
        [global::Microsoft.AspNetCore.Mvc.Razor.Internal.RazorInjectAttribute]
        public global::Microsoft.AspNetCore.Mvc.IUrlHelper Url { get; private set; }
        [global::Microsoft.AspNetCore.Mvc.Razor.Internal.RazorInjectAttribute]
        public global::Microsoft.AspNetCore.Mvc.IViewComponentHelper Component { get; private set; }
        [global::Microsoft.AspNetCore.Mvc.Razor.Internal.RazorInjectAttribute]
        public global::Microsoft.AspNetCore.Mvc.Rendering.IJsonHelper Json { get; private set; }
        [global::Microsoft.AspNetCore.Mvc.Razor.Internal.RazorInjectAttribute]
        public global::Microsoft.AspNetCore.Mvc.Rendering.IHtmlHelper<web_app_service.Models.SneakerProduct> Html { get; private set; }
    }
}
#pragma warning restore 1591
