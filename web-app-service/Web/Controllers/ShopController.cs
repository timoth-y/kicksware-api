﻿using System;
using System.Collections.Generic;
using System.Linq;
using Core.Services;
using Microsoft.AspNetCore.Mvc;
using SmartBreadcrumbs.Attributes;
using web_app_service.Models;
using web_app_service.Utils;

namespace web_app_service.Controllers
{
	public class ShopController : Controller
	{
		private readonly ISneakerProductService _service;

		public static readonly SneakerProductViewModel[] ProductsList =
		{
			new SneakerProductViewModel
			{
				UniqueId = "1",
				BrandName = "Nike",
				ModelName = "Air Max 720 Ispa",
				Description = "Crazy ones love them AF. My gf gave me these, I ll never sell them!",
				Price = 999m,
				ConditionIndex = 99.9m,
				Photos = {"Nike-ISPA-Air-Max-720.jpg"},
				AddedAt = new DateTime(2020, 2, 6)
			},
			new SneakerProductViewModel
			{
				UniqueId = "2",
				BrandName = "Nike",
				ModelName = "Jordan OG",
				Description = "Nike Jordan OG",
				Price = 200m,
				ConditionIndex = 89m,
				Photos = {"Nike_Jordan_OG.jpg"},
				AddedAt = new DateTime(2019, 12, 30)
			},
			new SneakerProductViewModel
			{
				UniqueId = "3",
				BrandName = "Nike",
				ModelName = "Air Force 1",
				Description = "Nike Air Force 1 Blue Wight Green",
				Price = 150m,
				ConditionIndex = 77m,
				Photos = {"Nike-Af1-Blue-Wight-Green.jpg"},
				AddedAt = new DateTime(2020, 2, 1)
			},
			new SneakerProductViewModel
			{
				UniqueId = "4",
				BrandName = "Nike",
				ModelName = "Jordan Proto Max",
				Description = "Nike Jordan Proto Max 720",
				Price = 240m,
				ConditionIndex = 100m,
				Photos = {"Nike-Jordan-Proto-Max-720.jpg"},
				AddedAt = new DateTime(2020, 1, 7)
			},
		};

		[ViewData]
		public string HeroCoverPath { get; set; } = "/images/heroes/shop-hero.jpg";

		[ViewData]
		public string HeroBreadTitle { get; set; } = "Buy sneakers";

		[ViewData]
		public string HeroBreadSubTitle { get; set; } = "Select and buy whatever kicks you like";

		public ShopController(ISneakerProductService service) => _service = service;

		[HttpGet]
		[Breadcrumb("Shop", FromAction = "Index", FromController = typeof(HomeController))]
		public IActionResult Products()
		{
			var products = _service.FetchAll()?.ToViewModel() ?? new List<SneakerProductViewModel>();

			return View(products);
		}

		[HttpGet]
		[Breadcrumb("Product item", FromAction = "Products", FromController = typeof(ShopController))]
		public IActionResult ProductItem(string productId)
		{
			var product = _service.FetchOne(productId)?.ToViewModel();
#if DEBUG
			if (product == null)
			{
				product = ProductsList.FirstOrDefault(p => p.UniqueId == productId);
			}
#endif
			if (product == null) return NotFound();
			ViewBag.RelatedProducts = ProductsList;
			return View(product);
		}
	}
}