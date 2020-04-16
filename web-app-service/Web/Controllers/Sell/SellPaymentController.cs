﻿using System.Linq;
using Microsoft.AspNetCore.Mvc;
using SmartBreadcrumbs.Attributes;
using web_app_service.Data.Reference_Data;
using web_app_service.Models;
using web_app_service.Wizards;

namespace web_app_service.Controllers
{
	public partial class SellController
	{
		[HttpGet]
		[Breadcrumb("Payment", FromAction = "NewProduct", FromController = typeof(SellController))]
		public ActionResult SetPayment(SneakerProductViewModel model)
		{
			HeroBreadSubTitle = "So let's talk about the price...";
			model.Price = 500m;
			AddBreadcrumbNode(nameof(Payment));
			return this.ViewStep(3, model);
		}

		[HttpPost]
		public ActionResult Payment(SneakerProductViewModel model, bool rollback)
		{
			if (rollback) SetPhotos(model);
			return SetShipping(model);
		}
	}
}