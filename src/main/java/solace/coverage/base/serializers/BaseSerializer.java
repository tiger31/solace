package solace.coverage.base.serializers;

import java.lang.reflect.Type;

import com.google.gson.JsonElement;
import com.google.gson.JsonObject;
import com.google.gson.JsonSerializationContext;

import solace.coverage.base.IMetric;
import solace.coverage.base.MetricMetadata;

public class BaseSerializer  {
	protected JsonElement serialize(IMetric<?> metric, Type metricType, JsonSerializationContext context) {
        JsonObject obj = new JsonObject();
        //Default metric fields
        obj.addProperty("type", metric.getClass().getSimpleName());
        obj.addProperty("coverage", metric.getCoverage());
        //Metric metadata
        MetricMetadata metadata = metric.getMetadata();
        obj.add("meta", (metadata.isEmpty()) ? null : context.serialize(metric.getMetadata()));
        //Inner metrics
        obj.add("inner", AnnotationsManager.getInnerMetricsFields(metric, context));
        return obj;
	}
}