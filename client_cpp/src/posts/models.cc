#include "models.hh"
#include "../../brunhild/mutations.hh"
#include "../state.hh"
#include "hide.hh"
#include "view.hh"
#include <sstream>

using nlohmann::json;
using std::string;

// Deserialize a property that might or might not be present from a kew of the
// same name
#define PARSE_OPT(key)                                                         \
    if (j.count(#key)) {                                                       \
        key = j[#key];                                                         \
    }

// Same as parse_opt, but explicitly converts to an std::string.
// Needed with std::optional<std::string> fields.
#define PARSE_OPT_STRING(key)                                                  \
    if (j.count(#key)) {                                                       \
        key = j.at(#key).get<string>();                                        \
    }

Image::Image(nlohmann::json& j)
{
    PARSE_OPT(audio);
    PARSE_OPT(video);
    PARSE_OPT(spoiler);
    file_type = static_cast<FileType>(j["fileType"]);
    thumb_type = static_cast<FileType>(j["thumbType"]);

    auto& j_dims = j["dims"];
    for (int i = 0; i < 4; i++) {
        dims[i] = j_dims[i];
    }

    PARSE_OPT(length);
    size = j["size"];
    PARSE_OPT_STRING(artist);
    PARSE_OPT_STRING(title);
    md5 = j["md5"];
    sha1 = j["sha1"];
    name = j["name"];
}

Command::Command(nlohmann::json& j)
{
    uint8_t _typ = j["type"];
    typ = static_cast<Type>(_typ);

    auto const& v = j["val"];
    switch (typ) {
    case Type::flip:
        val = v.get<bool>();
        break;
    case Type::eight_ball:
        eight_ball = v;
        break;
    case Type::pyu:
    case Type::pcount:
    case Type::rcount:
        val = v.get<unsigned long>();
        break;
    case Type::sync_watch: {
        std::array<unsigned long, 5> arr;
        for (int i = 0; i < 5; i++) {
            arr[i] = v[i];
        }
        val = arr;
    } break;
    case Type::dice: {
        std::array<uint16_t, 10> arr = { { 0 } };
        const int size = v.size();
        for (int i = 0; i < size; i++) {
            arr[i] = v[i];
        }
        val = arr;
    } break;
    case Type::roulette:
        val = std::array<uint8_t, 2>({ { v[0], v[1] } });
        break;
    }
}

string Image::image_root() const
{
    if (config.image_root_override != "") {
        return config.image_root_override;
    }
    return "/assets/images";
}

string Image::thumb_path() const
{
    std::ostringstream s;
    s << image_root() << "/thumb/" << sha1 << '.'
      << file_extentions.at(thumb_type);
    return s.str();
}

string Image::source_path() const
{
    std::ostringstream s;
    s << image_root() << "/src/" << sha1 << '.'
      << file_extentions.at(file_type);
    return s.str();
}

void Post::extend(nlohmann::json& j)
{
    PARSE_OPT(editing);
    PARSE_OPT(deleted);
    PARSE_OPT(sage);
    PARSE_OPT(banned);
    PARSE_OPT(sticky);
    PARSE_OPT(locked);

    id = j["id"];
    PARSE_OPT(op);
    time = j["time"];

    body = j["body"];
    PARSE_OPT(board);
    PARSE_OPT_STRING(name);
    PARSE_OPT_STRING(trip);
    PARSE_OPT_STRING(auth);
    PARSE_OPT_STRING(flag);

    if (j.count("image")) {
        image = Image(j["image"]);
    }
    parse_commands(j);
    parse_links(j);
}

void Post::parse_links(nlohmann::json& j)
{
    if (j.count("links")) {
        auto& l = j["links"];
        links.reserve(l.size());
        for (auto& val : l) {
            links[val["id"]] = { false, val["op"], val["board"] };
        }
    }
}

void Post::parse_commands(nlohmann::json& j)
{
    if (j.count("commands")) {
        auto& c = j["commands"];
        commands.clear(); // Not to duplicate existing entries
        commands.reserve(c.size());
        for (auto& com : c) {
            commands.push_back(Command(com));
        }
    }
}

void Post::propagate_links()
{

    // TODO: Notify about replies, if this post links to one of the user's posts

    for (auto&& [id, _] : links) {
        if (posts.count(id)) {
            auto& target = posts.at(id);
            target.backlinks[this->id] = LinkData { false, op, board };
            target.patch();
        }
        if (post_ids.hidden.count(id)) {
            hide_recursively(*this);
        }
    }
}

void Post::patch()
{
    for (auto& v : views) {
        v->patch();
    }
}

void Post::close()
{
    editing = false;
    patch();
}
