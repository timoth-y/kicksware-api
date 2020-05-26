﻿using System.Linq;
using Core.Entities.References;
using Microsoft.AspNetCore.Mvc;
using SmartBreadcrumbs.Attributes;

namespace Web.Controllers
{
	public partial class ShopController : Controller
	{
		[HttpGet]
		[Route("shop/model/{modelID}")]
		[Breadcrumb("Shop", FromAction = "Index", FromController = typeof(HomeController))]
		public IActionResult Model(string modelID, int page = 1, string sortBy = default)
		{
			var references = InitFilterHandler<SneakerReference>(new {modelID});
			if (!string.IsNullOrEmpty(sortBy)) references.ChooseSortParameter(sortBy);
			references.FetchPage(page);

			var brand = references.FirstOrDefault()?.Brand ?? new SneakerBrand(modelID);
			var baseModel = references.FirstOrDefault()?.Model ?? new SneakerModel(modelID);

			HeroCoverPath = baseModel.HeroPath ?? brand.HeroPath;
			HeroBreadTitle = baseModel.Name;
			HeroBreadSubTitle = brand.Description;
			HeroLogoPath = brand.HeroPath;

			AddBreadcrumbNode(nameof(Model), brand.Name);
			return View("References", references);
		}
	}
}