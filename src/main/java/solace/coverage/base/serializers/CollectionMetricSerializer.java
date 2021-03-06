package solace.coverage.base.serializers;

import java.lang.reflect.Type;

import com.google.gson.JsonElement;
import com.google.gson.JsonObject;
import com.google.gson.JsonSerializationContext;
import com.google.gson.JsonSerializer;
import solace.coverage.base.ICollectionMetric;

public class CollectionMetricSerializer extends BaseSerializer implements JsonSerializer<ICollectionMetric<?, ?>> {
    @Override
    public JsonElement serialize(ICollectionMetric<?, ?> metric, Type metricType, JsonSerializationContext context) {
        JsonObject obj = (JsonObject) super.serialize(metric, metricType, context);
        obj.add("metrics", context.serialize(metric.toArray()));
        return obj;
    }
}