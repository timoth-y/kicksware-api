﻿using RestSharp;

namespace Infrastructure.Gateway.REST.Search
{
	public class SearchRequest : RestRequest, IGatewayRestRequest
	{
		protected SearchRequest(string entity, string resource)
			: base($"search/{{entity}}{resource}", Method.GET)
		{
			AddParameter("entity", entity, ParameterType.UrlSegment);
		}
	}
}