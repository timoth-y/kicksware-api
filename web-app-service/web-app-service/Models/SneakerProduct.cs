﻿using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;

namespace web_app_service.Models
{
	public class SneakerProduct
	{
		public string Id { get; set; }
		public string Url { get; set; }
		public string BrandName { get; set; }
		public string ModelName { get; set; }
		[DataType(DataType.Currency)]
		public decimal Price { get; set; }
		public string Description { get; set; }
		public string Owner { get; set; }
		public List<string> Images => _images ??= new List<string>();
		public decimal StateIndex { get; set; }
		public DateTime AddedAt { get; set; }

		private List<string> _images = new List<string>();
	}
}
