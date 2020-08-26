package solace;

import io.swagger.v3.parser.*;
import solace.coverage.base.ICollectionMetric;
import solace.coverage.base.IMapMetric;
import solace.coverage.base.IMetric;
import solace.coverage.base.serializers.CollectionMetricSerializer;
import solace.coverage.base.serializers.MapMetricSerializer;
import solace.coverage.base.serializers.MetricSerializer;
import solace.coverage.metrics.PathsMetrics;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;

import io.swagger.v3.oas.models.OpenAPI;

public class App 
{
    public static void main( String[] args )
    {
        OpenAPI openAPI = new OpenAPIV3Parser().read("https://petstore3.swagger.io/api/v3/openapi.json");
        GsonBuilder builder = new GsonBuilder();
        PathsMetrics p = new PathsMetrics(openAPI.getPaths());
        builder.serializeNulls()
            .registerTypeHierarchyAdapter(IMetric.class, new MetricSerializer())
            .registerTypeHierarchyAdapter(ICollectionMetric.class, new CollectionMetricSerializer())
            .registerTypeHierarchyAdapter(IMapMetric.class, new MapMetricSerializer());
        Gson gson = builder.create();
        System.out.println(gson.toJson(p));
    }
}
